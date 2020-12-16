package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const cartTable = `CREATE TABLE IF NOT EXISTS cart(
						uid int NOT NULL,
						pd_id int NOT NULL,
						amount int NOT NULL,
						PRIMARY KEY(uid, pd_id),
						FOREIGN KEY(uid) REFERENCES user
						FOREIGN KEY(pd_id) REFERENCES product
					);`

// Cart store data of single product in user's cart
type Cart struct {
	UID    int
	PdID   int
	Amount int
}

// CartDB contain functions to use
type CartDB struct {
	insert     *sql.Stmt
	_delete    *sql.Stmt
	updateAmnt *sql.Stmt
	getAll     *sql.Stmt
	debug      *sql.Stmt
}

// CartDBInit prepare functions for database using
func CartDBInit(db *sql.DB) *CartDB {
	var err error
	cart := new(CartDB)

	cart.insert, err = db.Prepare("INSERT INTO cart VALUES(?,?,?);")
	if err != nil {
		panic(err)
	}

	cart._delete, err = db.Prepare("DELETE FROM cart WHERE uid=? AND pd_id=?;")
	if err != nil {
		panic(err)
	}

	cart.updateAmnt, err = db.Prepare("UPDATE cart SET amount=? WHERE uid=? AND pd_id=?;")
	if err != nil {
		panic(err)
	}

	cart.getAll, err = db.Prepare("SELECT * FROM product WHERE pd_id IN (SELECT pd_id FROM cart WHERE uid=?);")
	if err != nil {
		panic(err)
	}

	cart.debug, err = db.Prepare("SELECT * FROM cart;")
	if err != nil {
		panic(err)
	}

	return cart
}

// AddProductIntoCart add product into cart with pdid and amount
func (c *CartDB) AddProductIntoCart(uid, pdid, amount int) error {
	_, err := c.insert.Exec(uid, pdid, amount)
	return err
}

// DeleteProductFromCart delete product from cart with product id
func (c *CartDB) DeleteProductFromCart(id, pdid int) error {
	_, err := c._delete.Exec(id, pdid)
	return err
}

// UpdateAmount changes amount of product in cart of a user
// need to pass user id, product id and new amount
func (c *CartDB) UpdateAmount(uid, pdid, amount int) error {
	_, err := c.updateAmnt.Exec(amount)
	return err
}

// GetAllProductOfUser return all product id and amount by user id
func (c *CartDB) GetAllProductOfUser(uid int) (all []Product, totalPrice int) {
	rows, err := c.getAll.Query(uid)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var pd Product
		err = rows.Scan(&pd.Pdid, &pd.PdName, &pd.Price, &pd.Description, &pd.Amount, &pd.Eval, &pd.SellerID, &pd.Bid, &pd.Date)
		if err != nil {
			log.Println(err)
		}

		all = append(all, pd)
	}

	// this is a value that counts total price
	total := 0
	for i := range all {
		total += all[i].Price
	}

	return all, total
}

// Debug ing
func (c *CartDB) Debug() (all []Cart) {

	rows, err := c.debug.Query()
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var ca Cart
		err = rows.Scan(&ca.UID, &ca.PdID, &ca.Amount)
		if err != nil {
			log.Println(err)
			return
		}

		all = append(all, ca)
	}

	return
}
