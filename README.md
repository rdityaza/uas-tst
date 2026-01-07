# Encryption Microservice (Generic Domain)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

> **Tugas Besar Sistem Terintegrasi (II3160) - Tugas 2**
> Layanan backend yang menangani logika kriptografi (Security Context) menggunakan algoritma AES-GCM.

## 📖 Deskripsi
Layanan ini dirancang sebagai *Generic Domain* yang bersifat *stateless*. Fungsinya murni untuk mengamankan data teks (Plaintext) menjadi format terenkripsi (Ciphertext) dan sebaliknya. Layanan ini dibangun menggunakan bahasa **Go (Golang)** untuk performa komputasi tinggi.

## 🚀 Fitur Utama
* **AES-GCM Encryption:** Standar keamanan industri.
* **RESTful API:** Endpoint JSON untuk integrasi mudah.
* **Static File Serving:** Mampu menyajikan file frontend statis (embedded).

## 🛠️ Dokumentasi API
**Base URL:** `https://radit.tugastst.my.id`

| Method | Endpoint | Deskripsi | Request Body |
| :--- | :--- | :--- | :--- |
| `POST` | `/encrypt` | Mengubah teks menjadi sandi | `{"text": "...", "key": "..."}` |
| `POST` | `/decrypt` | Mengembalikan teks asli | `{"text": "...", "key": "..."}` |

### 🔑 Catatan Penting Pengujian (Key)

Algoritma AES-256 yang digunakan mewajibkan panjang kunci **TEPAT 32 Karakter (Bytes)**. Jika panjang kunci tidak sesuai (kurang atau lebih), API akan mengembalikan error. Key yang digunakan pada saat /decrypt harus SAMA PERSIS dengan Key yang digunakan saat /encrypt.

Untuk kemudahan pengujian, silakan gunakan **Default Key** berikut:

```text
12345678901234567890123456789012

## ⚙️ Cara Menjalankan (Docker)
```bash
# Build Image
docker build -t cipher-service .

# Run Container
docker run -d -p 8080:8080 cipher-service