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
	insert := false
	_, err := os.Stat(file)
	if err != nil {
		insert = true
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

	if insert {
		TestInsert()
	}
}

// RemoveAll : *FATAL* this command will remove whole database
func RemoveAll() {
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
