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
      git clone https://github.com/fdhlprtma/animal-api.git
      cd animal-api
   2. **Buat database & import schema**
      ```bash
      mysql -u root -p -e "CREATE DATABASE animaldb;"
      mysql -u root -p animaldb < internal/config/schema.sql
   3. **Salin file .env.example menjadi .env**
      ```bash
      cp .env.example .env
      ```
   4. **Isi variabel environment di .env**
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
  │   └── server/
  |       └── main.go          # Entry point aplikasi
  │
  ├── internal/
  │   ├── config/          # Konfigurasi database & env
  |   |   └── schema.sql   # Skema database
  │   ├── handler/         # Handler untuk endpoint API
  │   ├── model/           # Model data hewan
  │   ├── service/         # Logika bisnis
  |   ├── repository/      # Akses ke database & query SQL
  |   |   └── service/
  │   └── middleware/      # Middleware autentikasi JWT
  │
  ├── pkg/
  |   └── utils/
  |
  ├── uploads/             # Folder penyimpanan gambar
  ├── go.mod
  ├── go.sum
  ├── .env
  └── README.md
```
## Environment Variables (.env)
  ```graphql
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
      ```graphsql
      | Field                     | Tipe         | Contoh Value                                                    |
      | ------------------------- | ------------ | --------------------------------------------------------------- |
      | `name`                    | string       | `Rusa Merah`                                                    |
      | `image`                   | file         | *(pilih file JPG/PNG)*                                          |
      | `classification[kingdom]` | string       | `Animalia`                                                      |
      | `classification[phylum]`  | string       | `Chordata`                                                      |
      | `classification[class]`   | string       | `Mammalia`                                                      |
      | `classification[order]`   | string       | `Artiodactyla`                                                  |
      | `classification[family]`  | string       | `Cervidae`                                                      |
      | `characteristics`         | string       | `Tanduk hanya pada jantan, perut empat ruang.`                  |
      | `examples[]`              | string array | `Rusa Merah (Cervus elaphus)`, `Rusa Kutub (Rangifer tarandus)` |
      | `habitat`                 | string       | `Hutan, padang rumput, tundra`                                  |
      | `ecological_role`         | string       | `Mengendalikan vegetasi, mangsa predator`                       |
      ```
      ***Catatan :***
      - Gunakan form-data jika ingin mengunggah file gambar (image).
      - Gunakan raw JSON jika ingin menyertakan image_url langsung.
      - Struktur field lainnya sama persis.
   
   5. **Autentikasi (Login)**
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

## 🙏 Terima Kasih
   Terima kasih telah menggunakan Animal API!
   Jika kamu menemukan bug, punya saran fitur, atau ingin berkontribusi, jangan ragu untuk membuka issue atau pull request di repository ini.
   
   Selamat mencoba dan semoga API ini membantu dalam pengelolaan data hewan kamu! 🦌🚀
