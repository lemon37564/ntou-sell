package database

import (
	"database/sql"
	"log"
	"time"
)

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

	log.Println("All Table was successfully Created. Time:", time.Since(start))
}

func createUserTable(db *sql.DB) {
	_, err := db.Exec(userTable)
	logger("user", err)

	// insert one value into user in order to prevent max(uid) = null
	db.Exec("INSERT INTO user VALUES(0, \"N/A\", \"N/A\", \"superuser\", 0.0);")
}

func createProductTable(db *sql.DB) {
	_, err := db.Exec(productTable)
	logger("product", err)

	// insert one value into product in order to prevent max(pd_id) = null
	db.Exec("INSERT INTO product VALUES(0, \"N/A\", 0, \"N/A\", 0, 0.0, 0, false, \"2006-01-02\");")
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
	db.Exec("INSERT INTO history VALUES(0, 0, 0);")
}

func createOrderTable(db *sql.DB) {
	_, err := db.Exec(ordersTable)
	logger("orders", err)
}

func createMessageTable(db *sql.DB) {
	_, err := db.Exec(messageTable)
	logger("message", err)

	// insert one value into message in order to prevent max(mid) = null
	db.Exec("INSERT INTO history VALUES(0, 0, 0, \"none\");")
}

func logger(table string, err error) {
	if err != nil {
		log.Fatalf("Error Creating Table (%s) -> %v\n", table, err)
	}
	log.Printf("Table Created (%s)\n", table)
}
