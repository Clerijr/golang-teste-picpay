package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Initialize(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL!")
	return db, nil
}
