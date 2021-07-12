package database

import (
	"database/sql"
	"log"
	"os"

	// import go-sqlit3 for the sql driver
	_ "github.com/mattn/go-sqlite3"
)

const file = "database.db"

var DB *sql.DB

func init() {
	_, err := os.Stat(file)
	if err != nil {
		log.Println("no database file founded.")
	}

	DB, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	createTables(DB)

	bidPrepare(DB)
	cartPrepare(DB)
	historyPrepare(DB)
	messagePrepare(DB)
	orderPrepare(DB)
	productPrepare(DB)
	userPrepare(DB)

	TestInsert()
}

// RemoveAll : *FATAL* this command will remove whole database
func RemoveAll() {
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
