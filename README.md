# 🌀 Go-gin-fw

[![Go Version](https://img.shields.io/badge/Go-1.22+-blue.svg)](https://go.dev/)
[![Gin Framework](https://img.shields.io/badge/Gin-Framework-green.svg)](https://gin-gonic.com/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-success.svg)](#)

Aplikasi CRUD sederhana menggunakan **Go (Golang)** dengan framework **Gin** dan database **MySQL**.  
Project ini memiliki fitur **autentikasi (login & register)** serta **CRUD user** dengan template HTML.

---

## 🚀 Fitur Utama

- 🔐 Login dan Register pengguna  
- 👤 CRUD data user (Create, Read, Update, Delete)  
- 🧩 Routing modular (`routes/`)  

# 🌀 Go-gin-fw

[![Go Version](https://img.shields.io/badge/Go-1.22+-blue.svg)](https://go.dev/)
[![Gin Framework](https://img.shields.io/badge/Gin-Framework-green.svg)](https://gin-gonic.com/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-success.svg)](#)

Aplikasi CRUD sederhana menggunakan **Go (Golang)** dengan framework **Gin** dan database **MySQL**.  
Project ini memiliki fitur **autentikasi (login & register)** serta **CRUD user** dengan template HTML.

---

## 🚀 Fitur Utama

- 🔐 Login dan Register pengguna  
- 👤 CRUD data user (Create, Read, Update, Delete)  
- 🧩 Routing modular (`routes/`)  
- 🧱 Template HTML (`LoadHTMLGlob`)  
- 💾 Koneksi MySQL (`db/connect.go`)  
- 🛡️ Middleware autentikasi  

---

## 🗂️ Struktur Folder

Go-gin-fw/
├── main.go
├── database/
│ └── connect.go
├── routes/
│ └── routes.go
├── service/
│ ├── auth.go
│ ├── user.go
│ └── home.go
├── middleware/
│ └── auth.go
├── views/
│ ├── login.html
│ ├── register.html
│ └── user/
│ ├── users.html
│ ├── edit.html
│ └── confirm_delete.html
└── go.mod


---

## ⚙️ Setup Database

1. Buat database baru di MySQL:

   ```sql
   CREATE DATABASE go_gin_fw;
   USE go_gin_fw;

   CREATE TABLE users (
     id INT AUTO_INCREMENT PRIMARY KEY,
     username VARCHAR(100) NOT NULL,
     password VARCHAR(255) NOT NULL
   );

2. Atur koneksi MySQL di file database/connect.go:

const (
    DB_USER     = "root"
    DB_PASSWORD = ""
    DB_NAME     = "go_gin_fw"
)


🛠️ Cara Menjalankan

# Clone repo ini
git clone https://github.com/mhdfrxx/Go-gin-fw.git
cd Go-gin-fw

# Install dependency
go mod tidy

# Jalankan aplikasi
go run main.go

Lalu buka di browser:

http://localhost:8080/login

🧰 Teknologi yang Digunakan

| Komponen        | Teknologi     |
| --------------- | ------------- |
| Bahasa          | Go (Golang)   |
| Framework Web   | Gin           |
| Database        | MySQL         |
| Template Engine | html/template |

💡 Contoh Fitur
| Fitur               | Deskripsi                |
| ------------------- | ------------------------ |
| `/login`            | Halaman login user       |
| `/register`         | Registrasi pengguna baru |
| `/users`            | Menampilkan daftar user  |
| `/users/edit/:id`   | Edit data user           |
| `/users/delete/:id` | Konfirmasi & hapus user  |

🧑‍💻 Kontributor

Muhammad Farhan

📜 Lisensi

Proyek ini dilisensikan di bawah MIT License — bebas digunakan & dikembangkan lebih lanjut.
