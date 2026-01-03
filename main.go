package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// Struktur Tetap Sama
type Payload struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type Response struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

func setupCORS(w *http.ResponseWriter, r *http.Request) bool {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return true
	}
	return false
}

func main() {
	http.HandleFunc("/encrypt", encryptHandler)
	http.HandleFunc("/decrypt", decryptHandler)

	fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Service A (Cipher) with CORS running on port %s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("GAGAL START: ", err)
	}
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	if setupCORS(&w, r) {
		return 
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p Payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if len(p.Key) != 32 {
		p.Key = "12345678901234567890123456789012"
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
	if setupCORS(&w, r) {
		return
	}

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