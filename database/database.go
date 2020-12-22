package database

import (
	"database/sql"
	"log"
	"os"

	// import go-sqlit3 for the sql driver
	_ "github.com/mattn/go-sqlite3"
)

const file = "database.db"

// Data is a struct to manipulate database
type Data struct {
	bid     *bidStmt
	cart    *cartStmt
	history *historyStmt
	message *messageStmt
	order   *orderStmt
	product *productStmt
	user    *userStmt
}

// OpenAndInit open database and return *Data
func OpenAndInit() *Data {
	createTables()

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	data := Data{
		bid:     bidPrepare(db),
		cart:    cartPrepare(db),
		history: historyPrepare(db),
		message: messagePrepare(db),
		order:   orderPrepare(db),
		product: productPrepare(db),
		user:    userPrepare(db)}

	return &data
}

// RemoveAll : *FATAL* this command will remove whole database
func RemoveAll() {
	if err := os.RemoveAll(file); err != nil {
		log.Fatal(err)
	}
}
