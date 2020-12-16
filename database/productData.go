package database

import (
	"database/sql"
	"log"
	"time"
)

const productTable = `CREATE TABLE IF NOT EXISTS product(
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
						FOREIGN KEY(seller_uid) REFERENCES user
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

// ProductDB contain funcions to use
type ProductDB struct {
	insert        *sql.Stmt
	_delete       *sql.Stmt
	updatepdName  *sql.Stmt
	updatePrice   *sql.Stmt
	updateAmount  *sql.Stmt
	updateDecp    *sql.Stmt
	updateEval    *sql.Stmt
	maxpdID       *sql.Stmt
	newest        *sql.Stmt
	search        *sql.Stmt
	enhancesearch *sql.Stmt
	getPdInfo     *sql.Stmt
	allpd         *sql.Stmt
}

// ProductDBInit prepare function for database using
func ProductDBInit(db *sql.DB) *ProductDB {
	var err error
	product := new(ProductDB)

	product.insert, err = db.Prepare("INSERT INTO product VALUES(?,?,?,?,?,?,?,?,?);")
	if err != nil {
		panic(err)
	}

	product._delete, err = db.Prepare("DELETE FROM product WHERE seller_uid=? AND product_name=?;")
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

	product.updateDecp, err = db.Prepare("UPDATE product SET description=? WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	product.updateEval, err = db.Prepare("UPDATE product SET eval=? WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	product.maxpdID, err = db.Prepare("SELECT MAX(pd_id) FROM product;")
	if err != nil {
		panic(err)
	}

	product.newest, err = db.Prepare("SELECT * FROM product ORDER BY pd_id DESC LIMIT ?;")
	if err != nil {
		panic(err)
	}

	product.search, err = db.Prepare("SELECT * FROM product WHERE product_name LIKE ? AND pd_id>0 ORDER BY pd_id DESC;")
	if err != nil {
		panic(err)
	}

	product.enhancesearch, err = db.Prepare(`
		SELECT *
		FROM product
		WHERE product_name LIKE ? AND price>=? AND price<=? AND eval>=? AND pd_id>0
		ORDER BY pd_id DESC;
	`)
	if err != nil {
		panic(err)
	}

	product.getPdInfo, err = db.Prepare("SELECT * FROM product WHERE pd_id=? AND pd_id>0;")
	if err != nil {
		panic(err)
	}

	product.allpd, err = db.Prepare("SELECT * FROM product WHERE pd_id>0;")
	if err != nil {
		panic(err)
	}

	return product
}

// AddNewProduct add single product with product name, price, description, amount, seller id, bid and date into database
func (p *ProductDB) AddNewProduct(pdname string, price int, description string, amount int, sellerUID int, bid bool, date time.Time) (int, error) {
	var pdid int
	rows, err := p.maxpdID.Query()
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&pdid)
		if err != nil {
			log.Println(err)
			// panic if there's no product yet
		}
	}

	_, err = p.insert.Exec(pdid+1, pdname, price, description, amount, 0.0, sellerUID, bid, date)
	return pdid, err
}

// Delete product with product id
func (p *ProductDB) Delete(uid int, pdname string) error {
	_, err := p._delete.Exec(uid, pdname)

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
	if err != nil {
		log.Println(err)
		return err
	}
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
func (p *ProductDB) NewestProduct(number int) (all []Product) {
	rows, err := p.newest.Query(number)
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

// Search return product infos with searching keyword
func (p *ProductDB) Search(keyword string) (all []Product) {

	keyword = "%" + keyword + "%"

	rows, err := p.search.Query(keyword)
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

// SearchWithFilter is an enhance search function with filter
func (p *ProductDB) SearchWithFilter(keyword string, priceMin, priceMax, eval int) (all []Product) {

	keyword = "%" + keyword + "%"

	rows, err := p.enhancesearch.Query(keyword, priceMin, priceMax, eval)
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

// All return all products(debugging only)
func (p *ProductDB) All() (all []Product) {
	rows, err := p.allpd.Query()
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
