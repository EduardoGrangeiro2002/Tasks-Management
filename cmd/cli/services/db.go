package services

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Db () *sql.DB {
	db, err := sql.Open("sqlite3", "file:./db/db.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	return db
}