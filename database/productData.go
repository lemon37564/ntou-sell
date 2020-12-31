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
	FOREIGN KEY(seller_uid) REFERENCES user,
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

type productStmt struct {
	add     *sql.Stmt
	del     *sql.Stmt
	upName  *sql.Stmt
	upPrc   *sql.Stmt
	upAmt   *sql.Stmt
	upDes   *sql.Stmt
	upEval  *sql.Stmt
	maxPID  *sql.Stmt
	newest  *sql.Stmt
	search  *sql.Stmt
	filter  *sql.Stmt
	getInfo *sql.Stmt
	userPd  *sql.Stmt
}

func productPrepare(db *sql.DB) *productStmt {
	var err error
	product := new(productStmt)

	const (
		add    = "INSERT INTO product VALUES(?,?,?,?,?,?,?,?,?);"
		del    = "DELETE FROM product WHERE seller_uid=? AND product_name=?;"
		upName = "UPDATE product SET product_name=? WHERE pd_id=?;"
		upPrc  = "UPDATE product SET price=? WHERE pd_id=?;"
		upAmt  = "UPDATE product SET amount=? WHERE pd_id=?;"
		upDes  = "UPDATE product SET description=? WHERE pd_id=?;"
		upEval = "UPDATE product SET eval=? WHERE pd_id=?;"
		maxPID = "SELECT MAX(pd_id) FROM product;"
		newest = "SELECT * FROM product ORDER BY pd_id DESC LIMIT ?;"
		search = "SELECT * FROM product WHERE product_name LIKE ? AND pd_id>0 ORDER BY pd_id DESC;"
		filter = `
			SELECT *
			FROM product
			WHERE product_name LIKE ? AND price>=? AND price<=? AND eval>=? AND pd_id>0
			ORDER BY pd_id DESC;
			`
		getInfo = "SELECT * FROM product WHERE pd_id=? AND pd_id>0;"
		userPd  = "SELECT * FROM product WHERE seller_uid=?;"
	)

	if product.add, err = db.Prepare(add); err != nil {
		log.Println(err)
	}

	if product.del, err = db.Prepare(del); err != nil {
		log.Println(err)
	}

	if product.upName, err = db.Prepare(upName); err != nil {
		log.Println(err)
	}

	if product.upPrc, err = db.Prepare(upPrc); err != nil {
		log.Println(err)
	}

	if product.upAmt, err = db.Prepare(upAmt); err != nil {
		log.Println(err)
	}

	if product.upDes, err = db.Prepare(upDes); err != nil {
		log.Println(err)
	}

	if product.upEval, err = db.Prepare(upEval); err != nil {
		log.Println(err)
	}

	if product.maxPID, err = db.Prepare(maxPID); err != nil {
		log.Println(err)
	}

	if product.newest, err = db.Prepare(newest); err != nil {
		log.Println(err)
	}

	if product.search, err = db.Prepare(search); err != nil {
		log.Println(err)
	}

	if product.filter, err = db.Prepare(filter); err != nil {
		log.Println(err)
	}

	if product.getInfo, err = db.Prepare(getInfo); err != nil {
		log.Println(err)
	}

	if product.userPd, err = db.Prepare(userPd); err != nil {
		log.Println(err)
	}

	return product
}

// AddProduct add single product with product name, price, description, amount, seller id, bid and date into database
func (dt Data) AddProduct(name string, price int, description string, amount int, sellerUID int, bid bool, date time.Time) (int, error) {
	var pdid int
	rows, err := dt.Product.maxPID.Query()
	if err != nil {
		log.Println(err)
		return -1, err
	}

	for rows.Next() {
		err = rows.Scan(&pdid)
		if err != nil {
			log.Println(err)
			return -1, err
		}
	}
	pdid++

	_, err = dt.Product.add.Exec(pdid, name, price, description, amount, 0.0, sellerUID, bid, date)
	return pdid, err
}

// DeleteProduct with product id
func (dt Data) DeleteProduct(uid int, pdname string) error {
	_, err := dt.Product.del.Exec(uid, pdname)
	return err
}

// UpdateProductName with product id and new name
func (dt Data) UpdateProductName(pdid int, name string) error {
	_, err := dt.Product.upName.Exec(name, pdid)
	return err
}

// UpdateProductPrice with prouct id and new price
func (dt Data) UpdateProductPrice(pdid, price int) error {
	_, err := dt.Product.upPrc.Exec(price, pdid)
	return err
}

// UpdateProductAmount with prdouct id and new amount
func (dt Data) UpdateProductAmount(pdid, amount int) error {
	_, err := dt.Product.upAmt.Exec(amount, pdid)
	return err
}

// UpdateProductDescription with product id and new description
func (dt Data) UpdateProductDescription(pdid int, description string) error {
	_, err := dt.Product.upDes.Exec(description, pdid)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

// UpdateProductEval with product id and new eval
func (dt Data) UpdateProductEval(pdid int, eval float64) error {
	_, err := dt.Product.upEval.Exec(eval, pdid)
	return err
}

// GetProductInfoFromPdID return info of specific product with product id
func (dt Data) GetProductInfoFromPdID(pdid int) (pd Product) {
	rows, err := dt.Product.getInfo.Query(pdid)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&pd.Pdid, &pd.PdName, &pd.Price, &pd.Description, &pd.Amount, &pd.Eval, &pd.SellerID, &pd.Bid, &pd.Date)
		if err != nil {
			log.Println(err)
		}
	}

	return
}

// NewestProduct return newest number of products
func (dt Data) NewestProduct(number int) (all []Product) {
	rows, err := dt.Product.newest.Query(number)
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

	return
}

// SearchProduct return product infos with searching keyword
func (dt Data) SearchProduct(keyword string) (all []Product) {
	keyword = "%" + keyword + "%"

	rows, err := dt.Product.search.Query(keyword)
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

	return
}

// SearchProductWithFilter is an enhance search function with filter
func (dt Data) SearchProductWithFilter(keyword string, priceMin, priceMax, eval int) (all []Product) {
	keyword = "%" + keyword + "%"

	rows, err := dt.Product.filter.Query(keyword, priceMin, priceMax, eval)
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

	return
}

// GetSellerProduct list all product of a single seller
func (dt Data) GetSellerProduct(uid int) (all []Product) {
	rows, err := dt.Product.userPd.Query(uid)
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

	return
}
