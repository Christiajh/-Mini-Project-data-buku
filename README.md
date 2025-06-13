# 📚 Mini Project: Data Buku & Kategori (Go + Gin + PostgreSQL)

## 🚀 Deskripsi Project
Aplikasi RESTful API yang memungkinkan pengguna untuk melakukan CRUD data buku dan kategori. Proyek ini dibangun menggunakan:
- Golang + Gin
- PostgreSQL
- JSON Web Token (JWT) untuk autentikasi
- Dideploy ke Railway

---

## 🛠 Tech Stack
- Gin Web Framework (`github.com/gin-gonic/gin`)
- PostgreSQL driver (`github.com/lib/pq`)
- JWT Auth Middleware
- Railway for Cloud Deployment

---

## 📦 Struktur Database

### 📘 Tabel `users`
| Kolom       | Tipe Data | Keterangan          |
|-------------|-----------|---------------------|
| id          | serial    | primary key         |
| username    | varchar   |                     |
| password    | varchar   | hashed              |
| created_at  | timestamp | default now()       |

### 📚 Tabel `categories`
| Kolom       | Tipe Data | Keterangan          |
|-------------|-----------|---------------------|
| id          | serial    | primary key         |
| name        | varchar   | nama kategori       |
| created_at  | timestamp | default now()       |

### 📖 Tabel `books`
| Kolom         | Tipe Data | Keterangan                                         |
|---------------|-----------|----------------------------------------------------|
| id            | serial    | primary key                                       |
| title         | varchar   |                                                   |
| description   | text      |                                                   |
| image_url     | text      |                                                   |
| release_year  | integer   | validasi 1980 - 2024                              |
| price         | integer   | harga buku                                        |
| total_page    | integer   | total halaman                                     |
| thickness     | varchar   | otomatis 'tebal' atau 'tipis' dari total_page     |
| category_id   | integer   | foreign key ke kategori                           |
| created_at    | timestamp | default now()                                     |

---

