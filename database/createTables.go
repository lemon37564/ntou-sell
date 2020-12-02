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
	userTable := `
	CREATE TABLE user(
		id varchar(16) NOT NULL,
		account varchar(256) NOT NULL,
		password_hash varchar(64) NOT NULL,
		name varchar(256),
		eval float,             
		PRIMARY KEY(id)
	);
	`
	_, err := db.Exec(userTable)
	logger("user", err)
}

func createProductTable(db *sql.DB) {
	productTable := `
	CREATE TABLE product(
		pd_id varchar(16) NOT NULL,
		product_name varchar(256) NOT NULL,
		price int NOT NULL,
		description varchar(2048),
		amount int NOT NULL,
		eval float,
		id varchar(16) NOT NULL,
		bid bool,
		date varchar(16),
		PRIMARY KEY(pd_id),
		FOREIGN KEY(id) REFERENCES user
	);
	`
	_, err := db.Exec(productTable)
	logger("product", err)
}

func createBidTable(db *sql.DB) {
	bidTable := `
	CREATE TABLE bid(
		pd_id varchar(16) NOT NULL,
		deadline varchar(16),
		now_bidder_id varchar(16) NOT NULL,
		now_money int,
		seller_id varchar(16) NOT NULL,
		PRIMARY KEY(pd_id),
		FOREIGN KEY(seller_id) REFERENCES user
		FOREIGN KEY(now_bidder_id) REFERENCES user
	);
	`
	_, err := db.Exec(bidTable)
	logger("bid", err)
}

func createCartTable(db *sql.DB) {
	cartTable := `
	CREATE TABLE cart(
		id varchar(16) NOT NULL,
		products varchar(2048),
		amount int,
		PRIMARY KEY(id),
		FOREIGN KEY(id) REFERENCES user
	);
	`
	_, err := db.Exec(cartTable)
	logger("cart", err)
}

func createHistoryTable(db *sql.DB) {
	historyTable := `
	CREATE TABLE history(
		id varchar(16) NOT NULL,
		products varchar(2048),
		PRIMARY KEY(id),
		FOREIGN KEY(id) REFERENCES user
	);
	`
	_, err := db.Exec(historyTable)
	logger("history", err)
}

func createOrderTable(db *sql.DB) {
	// rename order as orders (order is a keword in SQL)
	orderTable := `
	CREATE TABLE orders(
		id varchar(16) NOT NULL,
		pd_id varchar(16) NOT NULL,
		name varchar(256),
		price int,
		amount int,
		sum int,
		seller_id varchar(16) NOT NULL,
		state varchar(8),
		PRIMARY KEY(id, pd_id),
		FOREIGN KEY(id) REFERENCES user,
		FOREIGN KEY(seller_id) REFERENCES user,
		FOREIGN KEY(pd_id) REFERENCES product
	);
	`
	_, err := db.Exec(orderTable)
	logger("orders", err)
}

func logger(table string, err error) {
	if err != nil {
		log.Fatalf("Error Creating Table (%s) -> %v\n", table, err)
	}
	log.Printf("Table Created (%s)\n", table)
}
