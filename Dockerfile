# Pakai Alpine polos
FROM alpine:latest

WORKDIR /root/

# 1. Copy file aplikasi 
COPY cipher-app .

# 2. Copy folder frontend 
COPY static ./static

# 3. Beri izin eksekusi
RUN chmod +x cipher-app

# 4. Jalankan
CMD ["./cipher-app"]