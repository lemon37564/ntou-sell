package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const cartTable = `CREATE TABLE cart(
	id varchar(16) NOT NULL,
	products varchar(2048),
	amount int,
	PRIMARY KEY(id),
	FOREIGN KEY(id) REFERENCES user
);`

type CartData struct {
	db *sql.DB

	insert     *sql.Stmt
	_delete    *sql.Stmt
	updatePds  *sql.Stmt
	updateAmnt *sql.Stmt
	_select    *sql.Stmt
}

func CartDataInit() *CartData {
	cart := new(CartData)

	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	cart.db = db

	insert, err := db.Prepare("INSERT INTO cart values(?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	cart.insert = insert

	_delete, err := db.Prepare("DELETE FROM cart where id=?;")
	if err != nil {
		log.Fatal(err)
	}
	cart._delete = _delete

	updatePds, err := db.Prepare("UPDATE cart SET products=?;")
	if err != nil {
		log.Fatal(err)
	}
	cart.updatePds = updatePds

	updateAmnt, err := db.Prepare("UPDATE cart SET amount=?;")
	if err != nil {
		log.Fatal(err)
	}
	cart.updateAmnt = updateAmnt

	_select, err := db.Prepare("SELECT ? FROM cart WHERE ?=?;")
	if err != nil {
		log.Fatal(err)
	}
	cart._select = _select

	return cart
}

func (c *CartData) Insert(id string, products string, amount int) error {
	_, err := c.insert.Exec(id, products, amount)
	return err
}

func (c *CartData) Delete(id string) error {
	_, err := c._delete.Exec(id)
	return err
}

func (c *CartData) UpdateProducts(products string) error {
	_, err := c.updatePds.Exec(products)
	return err
}

func (c *CartData) UpdateAmount(amount int) error {
	_, err := c.updateAmnt.Exec(amount)
	return err
}

// wait for implementation
func (c *CartData) Select() (string, error) {
	return "", nil
}

func (c *CartData) DBClose() error {
	return c.db.Close()
}
