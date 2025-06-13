package config

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	_ = godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Gagal koneksi DB: ", err)
	}
}