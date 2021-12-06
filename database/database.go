package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

const file = "database.db"

var db *sql.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	// retrieve the url
	dbURL := os.Getenv("DATABASE_URL")
	// connect to the db
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	createTables(db)

	bidPrepare(db)
	cartPrepare(db)
	historyPrepare(db)
	messagePrepare(db)
	orderPrepare(db)
	productPrepare(db)
	userPrepare(db)
	leaderBoardPrepare(db)

	TestInsert()
}

// RemoveAll : *FATAL* this command will remove whole database
func RemoveAll() {
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
