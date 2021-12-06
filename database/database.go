package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const file = "database.db"

var db *sql.DB

func init() {
	var err error
	// retrieve the url
	dbURL := os.Getenv("DATABASE_URL")
	// connect to the db
	db, err = sql.Open("postgres", dbURL)
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

func createTables(db *sql.DB) {
	log.Println("Initailizing database...")
	start := time.Now()

	createUserTable(db)
	createProductTable(db)
	createBidTable(db)
	createCartTable(db)
	createHistoryTable(db)
	createOrderTable(db)
	createMessageTable(db)
	_, err := db.Exec(leaderBoardTable)
	logger("leaderboard", err)

	log.Println("All Table was successfully Created. Time:", time.Since(start))
}

func createUserTable(db *sql.DB) {
	_, err := db.Exec(userTable)
	logger("user", err)

	// insert one value into user in order to prevent max(uid) = null
	_, err = userAdd.Exec(0, "test", "test", "superuser", 0.0)
	logger("insert first user", err)
}

func createProductTable(db *sql.DB) {
	_, err := db.Exec(productTable)
	logger("product", err)

	// insert one value into product in order to prevent max(pd_id) = null
	t, _ := time.Parse("2006-01-02", "2006-01-02")
	_, err = pdAdd.Exec(0, "test", 1, "test", 1, 0.0, 0, false, t)
	logger("insert first product", err)
}

func createBidTable(db *sql.DB) {
	_, err := db.Exec(bidTable)
	logger("bid", err)
}

func createCartTable(db *sql.DB) {
	_, err := db.Exec(cartTable)
	logger("cart", err)
}

func createHistoryTable(db *sql.DB) {
	_, err := db.Exec(historyTable)
	logger("history", err)

	// insert one value into history in order to prevent max(seq) = null
	_, err = histAdd.Exec(0, 0, 0)
	logger("insert first history", err)
}

func createOrderTable(db *sql.DB) {
	_, err := db.Exec(ordersTable)
	logger("orders", err)
}

func createMessageTable(db *sql.DB) {
	_, err := db.Exec(messageTable)
	logger("message", err)

	// insert one value into message in order to prevent max(mid) = null
	_, err = msgAdd.Exec(0, 0, 0, "null")
	logger("insert first message", err)
}

func logger(table string, err error) {
	if err != nil {
		log.Fatalf("Error Creating Table (%s) -> %v\n", table, err)
	}
	log.Printf("Table Created (%s)\n", table)
}
