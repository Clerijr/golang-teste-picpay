package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	dbURI := os.Getenv("POSTGRES_URL")
	db, err := sqlx.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database!")
	return db, nil
}
