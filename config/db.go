package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL tidak ditemukan! Pastikan environment variable diset.")
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Ping ke database gagal:", err)
	}

	log.Println("✅ Connected to database!")
}
