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
	Db *sql.DB

	Bid     *bidStmt
	Cart    *cartStmt
	History *historyStmt
	Message *messageStmt
	Order   *orderStmt
	Product *productStmt
	User    *userStmt
}

// OpenAndInit open database and return *Data
func OpenAndInit() *Data {
	createTables()

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	data := Data{
		Db:      db,
		Bid:     bidPrepare(db),
		Cart:    cartPrepare(db),
		History: historyPrepare(db),
		Message: messagePrepare(db),
		Order:   orderPrepare(db),
		Product: productPrepare(db),
		User:    userPrepare(db)}

	TestInsert(&data)

	return &data
}

// DBClose close the database file
func (d Data) DBClose() {
	d.Db.Close()
}

// RemoveAll : *FATAL* this command will remove whole database
func RemoveAll() {
	if err := os.RemoveAll(file); err != nil {
		log.Fatal(err)
	}
}
