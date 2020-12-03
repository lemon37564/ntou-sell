package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const productTable = `CREATE TABLE product(
						pd_id int NOT NULL,
						product_name varchar(256) NOT NULL,
						price int NOT NULL,
						description varchar(2048),
						amount int NOT NULL,
						eval float,
						seller_id int NOT NULL,
						bid bool,
						date sting,
						PRIMARY KEY(pd_id),
						FOREIGN KEY(seller_id) REFERENCES user
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
	Date        string
}

// ProductDB contain funcions to use
type ProductDB struct {
	insert       *sql.Stmt
	_delete      *sql.Stmt
	updatepdName *sql.Stmt
	updatePrice  *sql.Stmt
	updateAmount *sql.Stmt
	updateDecp   *sql.Stmt
	updateEval   *sql.Stmt
	maxpdID      *sql.Stmt
	search       *sql.Stmt
	getPdInfo    *sql.Stmt
}

// ProductDBInit prepare function for database using
func ProductDBInit(db *sql.DB) (product *ProductDB) {
	var err error

	product.insert, err = db.Prepare("INSERT INTO product VALUES(?,?,?,?,?,?,?,?,?);")
	if err != nil {
		panic(err)
	}

	product._delete, err = db.Prepare("DELETE FROM product WHERE pdid=?;")
	if err != nil {
		panic(err)
	}

	product.updatepdName, err = db.Prepare("UPDATE product SET product_name=? WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	product.updatePrice, err = db.Prepare("UPDATE product SET price=? WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	product.updateAmount, err = db.Prepare("UPDATE product SET amount=? WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	product.updateDecp, err = db.Prepare("UPDATE product SET decription=? WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	product.updateEval, err = db.Prepare("UPDARE product SET eval=? WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	product.maxpdID, err = db.Prepare("SELECT MAX(pd_id) FROM product;")
	if err != nil {
		panic(err)
	}

	product.search, err = db.Prepare("SELECT pd_id FROM product WHERE product_name LIKE%?%;")
	if err != nil {
		panic(err)
	}

	product.getPdInfo, err = db.Prepare("SELECT * FROM product WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	return product
}

// AddNewProduct add single product with product name, price, description, amount, seller id, bid and date into database
func (p *ProductDB) AddNewProduct(pdname string, price int, description string, amount int, sellerID int, bid bool, date string) error {
	var pdid int
	rows, err := p.maxpdID.Query()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&pdid)
		if err != nil {
			panic(err)
			// panic if there's no product yet
		}
	}

	_, err = p.insert.Exec(pdid+1, pdname, price, description, amount, 0.0, sellerID, bid, date)
	return err
}

// Delete product with product id
func (p *ProductDB) Delete(pdid int) error {
	_, err := p._delete.Exec(pdid)
	return err
}

// UpdateName with product id and new name
func (p *ProductDB) UpdateName(pdid int, name string) error {
	_, err := p.updatepdName.Exec(name, pdid)
	return err
}

// UpdatePrice with prouct id and new price
func (p *ProductDB) UpdatePrice(pdid, price int) error {
	_, err := p.updatePrice.Exec(price, pdid)
	return err
}

// UpdateAmount with prdouct id and new amount
func (p *ProductDB) UpdateAmount(pdid, amount int) error {
	_, err := p.updateAmount.Exec(amount, pdid)
	return err
}

// UpdateDescription with product id and new description
func (p *ProductDB) UpdateDescription(pdid int, description string) error {
	_, err := p.updateDecp.Exec(description, pdid)
	return err
}

// UpdateEval with product id and new eval
func (p *ProductDB) UpdateEval(pdid int, eval float64) error {
	_, err := p.updateEval.Exec(eval, pdid)
	return err
}

// GetInfoFromPdID return info of specific product with product id
func (p *ProductDB) GetInfoFromPdID(pdid int) (pd Product) {
	rows, err := p.getPdInfo.Query(pdid)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&pd.Pdid, &pd.PdName, &pd.Price, &pd.Description, &pd.Amount, &pd.Eval, &pd.SellerID, &pd.Bid, &pd.Date)
		if err != nil {
			panic(err)
		}
	}

	return
}

// Search return product ids with searching keyword
func (p *ProductDB) Search(keyword string) (pdid []int) {
	rows, err := p.search.Query(keyword)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var np int
		err = rows.Scan(&np)
		if err != nil {
			panic(err)
		}

		pdid = append(pdid, np)
	}

	return
}
