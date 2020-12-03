package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const cartTable = `CREATE TABLE cart(
	uid int NOT NULL,
	pd_id int NOT NULL,
	amount int NOT NULL,
	PRIMARY KEY(uid, pd_id),
	FOREIGN KEY(uid) REFERENCES user
	FOREIGN KEY(pd_id) REFERENCES product
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

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	cart.db = db

	cart.insert, err = db.Prepare("INSERT INTO cart VALUES(?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	cart._delete, err = db.Prepare("DELETE FROM cart WHERE uid=?;")
	if err != nil {
		log.Fatal(err)
	}

	cart.updatePds, err = db.Prepare("UPDATE cart SET products=?;")
	if err != nil {
		log.Fatal(err)
	}

	cart.updateAmnt, err = db.Prepare("UPDATE cart SET amount=?;")
	if err != nil {
		log.Fatal(err)
	}

	return cart
}

func (c *CartData) AddCart(id string, products string, amount int) error {
	_, err := c.insert.Exec(id, products, amount)
	return err
}

func (c *CartData) DeleteCart(id string) error {
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

// always use this function at the end
func (c *CartData) DBClose() error {
	return c.db.Close()
}
