package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const bidTable = `CREATE TABLE bid(
					pd_id varchar(16) NOT NULL,
					deadline varchar(16),
					now_bidder_id varchar(16) NOT NULL,
					now_money int,
					seller_id varchar(16) NOT NULL,
					PRIMARY KEY(pd_id),
					FOREIGN KEY(seller_id) REFERENCES user
					FOREIGN KEY(now_bidder_id) REFERENCES user
				);`

type BidData struct {
	db *sql.DB

	insert  *sql.Stmt
	_delete *sql.Stmt
	update  *sql.Stmt
	_select *sql.Stmt
}

func BidDataInit() *BidData {
	bid := new(BidData)

	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	bid.db = db

	insert, err := db.Prepare("INSERT INTO bid values(?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	bid.insert = insert

	_delete, err := db.Prepare("DELETE FROM bid where id=?;")
	if err != nil {
		log.Fatal(err)
	}
	bid._delete = _delete

	update, err := db.Prepare("UPDATE bid SET products=?;")
	if err != nil {
		log.Fatal(err)
	}
	bid.update = update

	_select, err := db.Prepare("SELECT ? FROM bid WHERE ?=?;")
	if err != nil {
		log.Fatal(err)
	}
	bid._select = _select

	return bid
}

// wait for implementation
func (b *BidData) Insert(id string, products string, amount int) error {
	_, err := b.insert.Exec(id, products, amount)
	return err
}

// wait for implementation
func (b *BidData) Delete(id string) error {
	_, err := b._delete.Exec(id)
	return err
}

// wait for implementation
func (b *BidData) UpdateProducts(products string) error {
	_, err := b.update.Exec(products)
	return err
}

// wait for implementation
func (b *BidData) Select() (string, error) {
	return "", nil
}

func (b *BidData) DBClose() error {
	return b.db.Close()
}
