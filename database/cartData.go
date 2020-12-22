package database

import (
	"database/sql"
	"log"
)

const cartTable = `
CREATE TABLE IF NOT EXISTS cart(
	uid int NOT NULL,
	pd_id int NOT NULL,
	amount int NOT NULL,
	PRIMARY KEY(uid, pd_id),
	FOREIGN KEY(uid) REFERENCES user,
	FOREIGN KEY(pd_id) REFERENCES product
);`

// Cart store data of single product in user's cart
type Cart struct {
	UID    int
	PdID   int
	Amount int
}

type cartStmt struct {
	add   *sql.Stmt
	del   *sql.Stmt
	upAmt *sql.Stmt
	get   *sql.Stmt
}

func cartPrepare(db *sql.DB) *cartStmt {
	var err error
	cart := new(cartStmt)

	const (
		add   = "INSERT INTO cart VALUES(?,?,?);"
		del   = "DELETE FROM cart WHERE uid=? AND pd_id=?;"
		upAmt = "UPDATE cart SET amount=? WHERE uid=? AND pd_id=?;"
		get   = "SELECT * FROM product WHERE pd_id IN (SELECT pd_id FROM cart WHERE uid=?);"
	)

	if cart.add, err = db.Prepare(add); err != nil {
		log.Println(err)
	}

	if cart.del, err = db.Prepare(del); err != nil {
		log.Println(err)
	}

	if cart.upAmt, err = db.Prepare(upAmt); err != nil {
		log.Println(err)
	}

	if cart.get, err = db.Prepare(get); err != nil {
		log.Println(err)
	}

	return cart
}

// AddProductIntoCart add product into cart with pdid and amount
func (dt Data) AddProductIntoCart(uid, pdid, amount int) error {
	_, err := dt.cart.add.Exec(uid, pdid, amount)
	return err
}

// DeleteProductFromCart delete product from cart with product id
func (dt Data) DeleteProductFromCart(id, pdid int) error {
	_, err := dt.cart.del.Exec(id, pdid)
	return err
}

// UpdateAmount changes amount of product in cart of a user
func (dt Data) UpdateAmount(uid, pdid, newAmount int) error {
	_, err := dt.cart.upAmt.Exec(newAmount, uid, pdid)
	return err
}

// GetAllProductOfUser return all product id and amount by user id
func (dt Data) GetAllProductOfUser(uid int) (all []Product, totalPrice int) {
	rows, err := dt.cart.get.Query(uid)
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
