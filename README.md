# Secure Messaging Platform - Encryption Microservice (Generic Domain)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Vue.js](https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D)

> **Tugas Besar Sistem Terintegrasi (II3160)**
> Microservice ini bertanggung jawab atas aspek keamanan (kriptografi) dalam sistem *Secure Messaging Platform*.

## 📖 Deskripsi Layanan

Layanan ini dirancang sebagai **Generic Domain** dalam arsitektur Domain-Driven Design (DDD) sistem kami. Berfungsi sebagai "brankas digital" yang terpisah dari logika komunikasi utama.

**Tanggung Jawab Utama:**
* Menangani **Security Context** sistem.
* Menyediakan API untuk mengubah *Plaintext* menjadi *Ciphertext* (Enkripsi).
* Menyediakan API untuk mengembalikan *Ciphertext* menjadi *Plaintext* (Dekripsi).
* Menyajikan antarmuka Frontend (Static Web) untuk demonstrasi integrasi.

Sistem menggunakan algoritma **AES-GCM (Advanced Encryption Standard - Galois/Counter Mode)** untuk menjamin kerahasiaan dan integritas data pesan.

## 🚀 Fitur Utama

1.  **Stateless Architecture:** Dibangun menggunakan **Go (Golang)** yang ringan dan cepat.
2.  **RESTful API:** Komunikasi standar menggunakan format JSON.
3.  **Cross-Origin Resource Sharing (CORS):** Mendukung integrasi dengan frontend dari domain berbeda.
4.  **Integrated Frontend:** Menyajikan UI Chatting berbasis Vue.js langsung dari server Go (Static File Embedding).
5.  **Dockerized:** Siap dideploy di mana saja (STB/Server) menggunakan Docker Container.

---

## 🛠️ Dokumentasi API (Tugas 2)

**Base URL:** `https://radit.tugastst.my.id`

### 1. Enkripsi Pesan
Mengubah teks biasa menjadi kode rahasia.

* **Endpoint:** `POST /encrypt`
* **Content-Type:** `application/json`
* **Request Body:**
    ```json
    {
      "text": "Halo, ini pesan rahasia",
      "key": "12345678901234567890123456789012"
    }
    ```
    *(Note: Key harus 32 karakter)*

* **Response (200 OK):**
    ```json
    {
      "result": "c29tZSBlbmNyeXB0ZWQgc3RyaW5n..."
    }
    ```

### 2. Dekripsi Pesan
Mengembalikan kode rahasia menjadi teks asli.

* **Endpoint:** `POST /decrypt`
* **Content-Type:** `application/json`
* **Request Body:**
    ```json
    {
      "text": "c29tZSBlbmNyeXB0ZWQgc3RyaW5n...",
      "key": "12345678901234567890123456789012"
    }
    ```

* **Response (200 OK):**
    ```json
    {
      "result": "Halo, ini pesan rahasia"
    }
    ```

---

## 💻 Integrasi Frontend (Tugas 3)

Layanan ini juga menyajikan antarmuka pengguna (Frontend) untuk mendemonstrasikan integrasi antara:
1.  **Identity & Communication Service** (Hakim)
2.  **Encryption Service** (Radit)

**Cara Akses:**
Buka Browser dan kunjungi `https://radit.tugastst.my.id/`

**Fitur Demo:**
* Registrasi & Login User (Integrasi API Auth Partner).
* Kirim Pesan Terenkripsi (Integrasi API Encrypt Lokal + API Chat Partner).
* Baca Pesan (Integrasi API History Partner + API Decrypt Lokal).
* Toggle View: Klik pesan untuk melihat bentuk asli vs terenkripsi.

---

## ⚙️ Cara Menjalankan (Deployment)

Project ini didesain untuk dijalankan menggunakan **Docker**.

### Prasyarat
* Docker & Docker Compose terinstall.
* Koneksi internet (untuk pull image Alpine).

### Langkah-langkah
1.  **Clone Repository:**
    ```bash
    git clone [https://github.com/rdityaza/uas-tst.git](https://github.com/rdityaza/uas-tst.git)
    cd uas-tst
    ```

2.  **Build Docker Image:**
    ```bash
    docker build -t cipher-service .
    ```

3.  **Jalankan Container:**
    ```bash
    docker run -d -p 8082:8080 --name cipher-service --restart=always cipher-service
    ```

4.  **Verifikasi:**
    Layanan berjalan di port `8082`.
    * API: `http://localhost:8082/encrypt`
    * Web: `http://localhost:8082/`


*Project ini dibuat untuk memenuhi Tugas Besar Mata Kuliah II3160 Teknologi Sistem Terintegrasi, ITB 2026.*