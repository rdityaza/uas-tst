# Kita pakai Alpine polos (kecil banget, cuma 5MB)
FROM alpine:latest

WORKDIR /root/

# Copy file matang dari laptop ke dalam image
COPY cipher-app .

# Kasih izin biar bisa dijalankan
RUN chmod +x cipher-app

# Jalankan aplikasinya
CMD ["./cipher-app"]