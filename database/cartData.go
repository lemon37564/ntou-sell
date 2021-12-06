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
	FOREIGN KEY(uid) REFERENCES userDB ON DELETE CASCADE,
	FOREIGN KEY(pd_id) REFERENCES product ON DELETE CASCADE,
	CHECK (amount > 0)
);`

// Cart store data of single product in user's cart
type Cart struct {
	UID    int
	PdID   int
	Amount int
}

var (
	cartAdd   *sql.Stmt
	cartDel   *sql.Stmt
	cartUpAmt *sql.Stmt
	cartGet   *sql.Stmt
)

func cartPrepare(db *sql.DB) {
	var err error

	const (
		add   = "INSERT INTO cart VALUES($1,$2,$3);"
		del   = "DELETE FROM cart WHERE uid=$1 AND pd_id=$2;"
		upAmt = "UPDATE cart SET amount=$1 WHERE uid=$2 AND pd_id=$3;"
		get   = "SELECT * FROM product WHERE pd_id IN (SELECT pd_id FROM cart WHERE uid=$1);"
	)

	if cartAdd, err = db.Prepare(add); err != nil {
		panic(err)
	}

	if cartDel, err = db.Prepare(del); err != nil {
		panic(err)
	}

	if cartUpAmt, err = db.Prepare(upAmt); err != nil {
		panic(err)
	}

	if cartGet, err = db.Prepare(get); err != nil {
		panic(err)
	}
}

// AddProductIntoCart add product into cart with pdid and amount
func AddProductIntoCart(uid, pdid, amount int) error {
	_, err := cartAdd.Exec(uid, pdid, amount)
	return err
}

// DeleteProductFromCart delete product from cart with product id
func DeleteProductFromCart(id, pdid int) error {
	_, err := cartDel.Exec(id, pdid)
	return err
}

// UpdateCartAmount changes amount of product in cart of a user
func UpdateCartAmount(uid, pdid, newAmount int) error {
	_, err := cartUpAmt.Exec(newAmount, uid, pdid)
	return err
}

// GetAllProductOfUser return all product id and amount by user id
func GetAllProductOfUser(uid int) (all []Product, totalPrice int) {
	rows, err := cartGet.Query(uid)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
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
