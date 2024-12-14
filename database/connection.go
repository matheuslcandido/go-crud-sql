package database

import (
	"database/sql"
	"log"
)

func Connect(logger *log.Logger) (*sql.DB, error) {
	logger.Print("Starting connect on DB!")

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	logger.Print("DB Connected!")

	return db, nil
}
