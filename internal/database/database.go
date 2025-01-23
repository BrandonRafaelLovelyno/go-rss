package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect(pgUrl string) *Queries {
	db, err := sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal("Error making pool connection: ", err)
	}

	return &Queries{db: db}
}
