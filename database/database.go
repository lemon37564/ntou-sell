package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	root = "C:/software-engineering"
	name = "database.db"
	file = root + "/" + name
)

// Open database and return *sql.DB
func Open() *sql.DB {
	check()

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}

	return db
}

// check if there's database exists
// if no, init.
func check() {
	_, err := os.Stat(root)
	if err != nil {
		createFolder()
	}

	_, err = os.Stat(file)
	if err != nil {
		createTables()
	}
}

func createFolder() {
	err := os.Mkdir(root, 0666)
	if err != nil {
		panic(err)
	}
}

// RemoveAll : *FATAL* this command will remove whole database
func RemoveAll() {
	err := os.RemoveAll(root)
	if err != nil {
		panic(err)
	}
}
