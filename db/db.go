package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Initialize(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
