package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(dsn string) {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
