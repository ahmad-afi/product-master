# Test Superindo (Product Master)

Cara Menjalankan program

1. Pastikan [docker](https://www.docker.com/) sudah terinstall
2. Copy `.env-exampple` menjadi `.env` lalu isi variabel yang sesuai
3. Jalankan perintah `docker compose up -d`
4. Jalankan migrtions dan seeder dengan `go run migrations/main.go -steps=1 -forceMigration=false` atau dengan `make migrate-up`
5. Import collection dari `https://documenter.getpostman.com/view/20402111/2sAYkEqzMu`

Untuk proses migrations, bisa dilihat readme didalam folder migrations
