package database

import (
	"database/sql"
	"log"
	"time"
)

const productTable = `
CREATE TABLE IF NOT EXISTS product(
	pd_id int NOT NULL,
	product_name varchar(64) NOT NULL,
	price int NOT NULL,
	description varchar(256),
	amount int NOT NULL,
	eval float,
	seller_uid int NOT NULL,
	bid bool,
	date timestamp,
	PRIMARY KEY(pd_id),
	FOREIGN KEY(seller_uid) REFERENCES userDB ON DELETE CASCADE,
	CHECK (price > 0 AND amount > 0)
);`

// Product type store data of a single product
type Product struct {
	Pdid        int
	PdName      string
	Price       int
	Description string
	Amount      int
	Eval        float64
	SellerID    int
	Bid         bool
	Date        time.Time
}

var (
	pdAdd     *sql.Stmt
	pdDel     *sql.Stmt
	pdUpName  *sql.Stmt
	pdUpPrc   *sql.Stmt
	pdUpAmt   *sql.Stmt
	pdUpDes   *sql.Stmt
	pdUpEval  *sql.Stmt
	pdMaxPID  *sql.Stmt
	pdNewest  *sql.Stmt
	pdSearch  *sql.Stmt
	pdFilter  *sql.Stmt
	pdGetInfo *sql.Stmt
	pdUserPd  *sql.Stmt
)

func productPrepare(db *sql.DB) {
	var err error

	const (
		add    = "INSERT INTO product VALUES($1,$2,$3,$4,$5,$6,$7,%8,$9);"
		del    = "DELETE FROM product WHERE seller_uid=$1 AND pd_id=$2;"
		upName = "UPDATE product SET product_name=$1 WHERE pd_id=$2;"
		upPrc  = "UPDATE product SET price=$1 WHERE pd_id=$2;"
		upAmt  = "UPDATE product SET amount=$1 WHERE pd_id=$2;"
		upDes  = "UPDATE product SET description=$1 WHERE pd_id=$2;"
		upEval = "UPDATE product SET eval=$1 WHERE pd_id=$2;"
		maxPID = "SELECT MAX(pd_id) FROM product;"
		newest = "SELECT * FROM product ORDER BY pd_id DESC LIMIT $1;"
		search = "SELECT * FROM product WHERE product_name LIKE $1 AND pd_id>0 ORDER BY pd_id DESC;"
		filter = `
			SELECT *
			FROM product
			WHERE product_name LIKE $1 AND price>=$2 AND price<=$3 AND eval>=$4 AND pd_id>0
			ORDER BY pd_id DESC;
			`
		getInfo = "SELECT * FROM product WHERE pd_id=$1 AND pd_id>0;"
		userPd  = "SELECT * FROM product WHERE seller_uid=$1 AND pd_id>0;"
	)

	if pdAdd, err = db.Prepare(add); err != nil {
		panic(err)
	}

	if pdDel, err = db.Prepare(del); err != nil {
		panic(err)
	}

	if pdUpName, err = db.Prepare(upName); err != nil {
		panic(err)
	}

	if pdUpPrc, err = db.Prepare(upPrc); err != nil {
		panic(err)
	}

	if pdUpAmt, err = db.Prepare(upAmt); err != nil {
		panic(err)
	}

	if pdUpDes, err = db.Prepare(upDes); err != nil {
		panic(err)
	}

	if pdUpEval, err = db.Prepare(upEval); err != nil {
		panic(err)
	}

	if pdMaxPID, err = db.Prepare(maxPID); err != nil {
		panic(err)
	}

	if pdNewest, err = db.Prepare(newest); err != nil {
		panic(err)
	}

	if pdSearch, err = db.Prepare(search); err != nil {
		panic(err)
	}

	if pdFilter, err = db.Prepare(filter); err != nil {
		panic(err)
	}

	if pdGetInfo, err = db.Prepare(getInfo); err != nil {
		panic(err)
	}

	if pdUserPd, err = db.Prepare(userPd); err != nil {
		panic(err)
	}
}

// AddProduct add single product with product name, price, description, amount, seller id, bid and date into database
func AddProduct(name string, price int, description string, amount int, sellerUID int, bid bool, date time.Time) (int, error) {
	var pdid int
	rows, err := pdMaxPID.Query()
	if err != nil {
		log.Println(err)
		return -1, err
	}

	if rows.Next() {
		err = rows.Scan(&pdid)
		if err != nil {
			log.Println(err)
			return -1, err
		}
	}

	rows.Close()
	_, err = pdAdd.Exec(pdid+1, name, price, description, amount, 0.0, sellerUID, bid, date)
	return pdid, err
}

// DeleteProduct with product id
func DeleteProduct(uid int, pdid int) error {
	_, err := pdDel.Exec(uid, pdid)
	return err
}

// UpdateProductName with product id and new name
func UpdateProductName(pdid int, name string) error {
	_, err := pdUpName.Exec(name, pdid)
	return err
}

// UpdateProductPrice with prouct id and new price
func UpdateProductPrice(pdid, price int) error {
	_, err := pdUpPrc.Exec(price, pdid)
	return err
}

// UpdateProductAmount with prdouct id and new amount
func UpdateProductAmount(pdid, amount int) error {
	_, err := pdUpAmt.Exec(amount, pdid)
	return err
}

// UpdateProductDescription with product id and new description
func UpdateProductDescription(pdid int, description string) error {
	_, err := pdUpDes.Exec(description, pdid)
	return err
}

// UpdateProductEval with product id and new eval
func UpdateProductEval(pdid int, eval float64) error {
	_, err := pdUpEval.Exec(eval, pdid)
	return err
}

// GetProductInfoFromPdID return info of specific product with product id
func GetProductInfoFromPdID(pdid int) (pd Product) {
	rows, err := pdGetInfo.Query(pdid)
	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&pd.Pdid, &pd.PdName, &pd.Price, &pd.Description, &pd.Amount, &pd.Eval, &pd.SellerID, &pd.Bid, &pd.Date)
		if err != nil {
			log.Println(err)
		}
	}

	return
}

// NewestProduct return newest number of products
func NewestProduct(number int) (all []Product) {
	rows, err := pdNewest.Query(number)
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

	return
}

// SearchProduct return product infos with searching keyword
func SearchProduct(keyword string) (all []Product) {
	keyword = "%" + keyword + "%"

	rows, err := pdSearch.Query(keyword)
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

	return
}

// SearchProductWithFilter is an enhance search function with filter
func SearchProductWithFilter(keyword string, priceMin, priceMax, eval int) (all []Product) {
	keyword = "%" + keyword + "%"

	rows, err := pdFilter.Query(keyword, priceMin, priceMax, eval)
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

	return
}

// GetSellerProduct list all product of a single seller
func GetSellerProduct(uid int) (all []Product) {
	rows, err := pdUserPd.Query(uid)
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

	return
}
