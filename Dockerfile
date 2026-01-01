# Stage 1: Build
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Stage 2: Run (Sangat ringan, cuma pakai Alpine)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080
CMD ["./main"]