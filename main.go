package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// Struktur Request & Response
type Payload struct {
	Text string `json:"text"`
	Key  string `json:"key"` // Key rahasia (harus 32 karakter untuk AES-256)
}

type Response struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

func main() {
	http.HandleFunc("/encrypt", encryptHandler)
	http.HandleFunc("/decrypt", decryptHandler)

	// Jalan di port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	println("Service A (Cipher) running on port " + port + "...")
	http.ListenAndServe(":"+port, nil)
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p Payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Pastikan Key 32 byte (bisa di-hardcode atau dari user)
	// Agar simpel buat tugas, jika user tidak kirim key, kita pakai default
	if len(p.Key) != 32 {
		p.Key = "12345678901234567890123456789012" // Default Key 32 chars
	}

	block, _ := aes.NewCipher([]byte(p.Key))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	
	ciphertext := gcm.Seal(nonce, nonce, []byte(p.Text), nil)
	result := base64.StdEncoding.EncodeToString(ciphertext)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Result: result})
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p Payload
	json.NewDecoder(r.Body).Decode(&p)

	if len(p.Key) != 32 {
		p.Key = "12345678901234567890123456789012"
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(p.Text)
	block, _ := aes.NewCipher([]byte(p.Key))
	gcm, _ := cipher.NewGCM(block)
	
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		json.NewEncoder(w).Encode(Response{Error: "Ciphertext too short"})
		return
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	
	if err != nil {
		json.NewEncoder(w).Encode(Response{Error: "Decryption failed"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Result: string(plaintext)})
}