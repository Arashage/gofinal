package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() (err error) {
	database := os.Getenv("DATABASE_URL")
	DB, err = sql.Open("postgres", database)
	return
}
