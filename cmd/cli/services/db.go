package services

import (
	"database/sql"
	"log"
	"os"
	"tasks/utils"

	_ "modernc.org/sqlite"
)


func Db () *sql.DB {
	path := utils.DbPath()
	db, err := sql.Open("sqlite", "file:"+path+"db.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateDirIfNotExist() {
	path := utils.DbPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Fatalf("Error creating directory: %v", err)
		}
	}
}