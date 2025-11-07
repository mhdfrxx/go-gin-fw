package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_products")
	if err != nil {
		log.Fatal("❌ Gagal konek ke MySQL:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("❌ Database tidak bisa diakses:", err)
	}
	log.Println("✅ Terhubung ke database MySQL")
}
