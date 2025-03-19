# Test Superindo (Product Master)

Cara Menjalankan program

1. Pastikan [docker](https://www.docker.com/) sudah terinstall
2. Jalankan perintah `docker compose up redis_product_master postgre_product_master -d`
3. Import collection dari `https://documenter.getpostman.com/view/20402111/2sAYkEqzMu`
4. Copy `.env-exampple` menjadi `.env` lalu isi variabel yang sesuai
5. Jalankan program dengan `go run main.go`
