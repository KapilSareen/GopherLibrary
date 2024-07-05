package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

var db *sql.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg := mysql.Config{
		User:   os.Getenv("DATABASE_USERNAME"),
		Passwd: os.Getenv("DATABASE_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("URL"),
		DBName: os.Getenv("DATABASE_NAME"),
	}

	db, err = sql.Open("mysql", strings.Split(cfg.FormatDSN(), "?")[0])
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}
