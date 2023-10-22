package bot

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initializeDatabse(path string) (db *sql.DB, err error) {
	log.Printf("Initializing database: %s", path)
	db, err = sql.Open("sqlite3", fmt.Sprintf("file:%s", path))
	if err != nil {
		return db, err
	}

	_, err = db.Exec("Create table if not exists news (" +
		"guid VARCHAR(30) PRIMARY KEY" +
		")")

	if err != nil {
		return db, err
	}

	log.Printf("Initialized database: %s", path)
	return db, nil
}
