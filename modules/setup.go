package modules

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ConnectDatabase() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			user_id INTEGER PRIMARY KEY,
			first_name TEXT,
			last_name TEXT,
			username TEXT,
			created_at TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

}
