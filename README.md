# 🦌 Animal API

API profesional berbasis Golang untuk mengelola data hewan, dengan dukungan autentikasi JWT, upload gambar, pencarian, dan dokumentasi endpoint yang jelas.  
Data hewan mencakup informasi biologis, karakteristik, contoh, habitat, dan peran ekologis.

---

## ✨ Fitur

- **GET publik** untuk semua data hewan
- **GET by ID** untuk detail hewan tertentu (`/{id}`)
- **Pencarian** hewan berdasarkan nama (`?q=nama`)
- **Upload gambar** lewat `multipart/form-data`
- **Autentikasi JWT** untuk endpoint `POST` (buat data baru)
- **Struktur proyek profesional** (modular & maintainable)
- **Database MySQL** dengan skema siap pakai (`schema.sql`)

---

## 📦 Persyaratan

Sebelum menjalankan API, pastikan sudah menginstal:

- [Go](https://go.dev/dl/) minimal versi **1.22**
- [MySQL](https://dev.mysql.com/downloads/)
- [Git](https://git-scm.com/downloads)

---

## ⚙️ Instalasi & Setup

1. **Clone repositori**
   ```bash
   git clone https://github.com/username/animal-api.git
   cd animal-api
2. **Buat database & import schema**
   ```bash
   mysql -u root -p -e "CREATE DATABASE animaldb;"
   mysql -u root -p animaldb < internal/config/schema.sql
3. **Isi variabel environment di .env**
   ```bash
   DB_USER=root
   DB_PASSWORD=
   DB_NAME=animaldb
   DB_HOST=localhost
   DB_PORT=3306
   JWT_SECRET=your_jwt_secret_key

## Struktur Folder
  ```graphql
  animal-api/
  │
  ├── cmd/
  │   └── main.go          # Entry point aplikasi
  │
  ├── internal/
  │   ├── config/          # Konfigurasi database & env
  |   |   └── schema.sql   # Skema database
  │   ├── handler/         # Handler untuk endpoint API
  │   ├── model/           # Model data hewan
  │   ├── service/         # Logika bisnis
  │   └── middleware/      # Middleware autentikasi JWT
  │
  ├── uploads/             # Folder penyimpanan gambar
  ├── go.mod
  ├── go.sum
  ├── .env
  └── README.md
```
## Environment Variables (.env)
  ```bash
  | Variabel     | Deskripsi               | Contoh     |
  | ------------ | ----------------------- | ---------- |
  | DB\_USER     | Username database MySQL | root       |
  | DB\_PASSWORD | Password database MySQL | (kosong)   |
  | DB\_NAME     | Nama database           | animaldb   |
  | DB\_HOST     | Host database           | localhost  |
  | DB\_PORT     | Port database           | 3306       |
  | JWT\_SECRET  | Secret key untuk JWT    | rahasia123 |
```

## Menjalankan API
  ```bash
    go run cmd/main.go
  ```
  **Atau build terlebih dahulu:**
  ```bash
    go build -o animal-api cmd/main.go
    ./animal-api
  ```
  **API akan berjalan di:**
  ```bash
    http://localhost:8080
  ```

## Endpoint 
1. **GET Semua Hewan**
   ```bash
   GET /animals
   ```
   ***contoh :***
   ```hash
   GET /animals?q=harimau
2. **GET Hewan Berdasarkan ID**
   ```bash
   GET /animals/{id}
3. **POST Tambah Hewan (Autentikasi JWT)**
   ```bash
   POST /animals
   Content-Type: multipart/form-data
   Authorization: Bearer <token>
   ```
   ***Field:***
    - name (string)
    - image (file)
    - classification[kingdom] ... classification[species]
    - characteristics (string)
    - examples[] (array string)
    - habitat (string)
    - ecological_role (string)
4. **Autentikasi (Login)**
   ```bash
   POST /login
   Content-Type: application/json
   ```
   ***Body:***
   ```bash
   {
    "username": "admin",
    "password": "admin123"
   }
   ```
