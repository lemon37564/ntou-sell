package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

const file = "database.db"

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	//透過 Getenv 來讀取 .env
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	//連結 db
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	//錯誤攔截與建立連接
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	createTables(db.DB())

	bidPrepare(db.DB())
	cartPrepare(db.DB())
	historyPrepare(db.DB())
	messagePrepare(db.DB())
	orderPrepare(db.DB())
	productPrepare(db.DB())
	userPrepare(db.DB())
	leaderBoardPrepare(db.DB())

	TestInsert()
}

// RemoveAll : *FATAL* this command will remove whole database
func RemoveAll() {
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
