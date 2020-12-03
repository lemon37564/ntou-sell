package database

import (
	"database/sql"
	"log"
	"time"
)

func createTables() {
	log.Println("Initailizing database...")
	start := time.Now()

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createUserTable(db)
	createProductTable(db)
	createBidTable(db)
	createCartTable(db)
	createHistoryTable(db)
	createOrderTable(db)

	log.Println("All Table was successfully Created. Time:", time.Since(start))
}

func createUserTable(db *sql.DB) {
	_, err := db.Exec(userTable)
	tableLogger("user", err)
}

func createProductTable(db *sql.DB) {
	_, err := db.Exec(productTable)
	tableLogger("product", err)
}

func createBidTable(db *sql.DB) {
	_, err := db.Exec(bidTable)
	tableLogger("bid", err)
}

func createCartTable(db *sql.DB) {
	_, err := db.Exec(cartTable)
	tableLogger("cart", err)
}

func createHistoryTable(db *sql.DB) {
	_, err := db.Exec(historyTable)
	tableLogger("history", err)
}

func createOrderTable(db *sql.DB) {
	_, err := db.Exec(ordersTable)
	tableLogger("orders", err)
}

func tableLogger(table string, err error) {
	if err != nil {
		log.Fatalf("Error Creating Table (%s) -> %v\n", table, err)
	}
	log.Printf("Table Created (%s)\n", table)
}
