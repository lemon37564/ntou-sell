package database

import (
	"database/sql"
	"os"

	// import go-sqlit3 for the sql driver
	_ "github.com/mattn/go-sqlite3"
)

const (
	file = "database.db"
)

// Open database and return *sql.DB
func Open() *sql.DB {
	createTables()

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}

	return db
}

// RemoveAll : *FATAL* this command will remove whole database
func RemoveAll() {
	err := os.RemoveAll(file)
	if err != nil {
		panic(err)
	}
}
