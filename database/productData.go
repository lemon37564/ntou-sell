package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const productTable = `CREATE TABLE product(
						pd_id int NOT NULL,
						product_name varchar(256) NOT NULL,
						price int NOT NULL,
						description varchar(2048),
						amount int NOT NULL,
						eval float,
						uid int NOT NULL,
						bid bool,
						date varchar(16),
						PRIMARY KEY(pd_id),
						FOREIGN KEY(uid) REFERENCES user
					);`

type ProductData struct {
	db *sql.DB

	insert       *sql.Stmt
	_delete      *sql.Stmt
	updatePrice  *sql.Stmt
	updateAmount *sql.Stmt
}

func ProductDataInit() *ProductData {
	product := new(ProductData)

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	product.db = db

	product.insert, err = db.Prepare("INSERT INTO product VALUES(?,?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	product._delete, err = db.Prepare("DELETE FROM product WHERE pdid=?;")
	if err != nil {
		log.Fatal(err)
	}

	product.updatePrice, err = db.Prepare("UPDATE product SET price=? WHERE pd_id=?;")
	if err != nil {
		log.Fatal(err)
	}

	product.updateAmount, err = db.Prepare("UPDATE product SET amount=? WHERE pd_id=?;")
	if err != nil {
		log.Fatal(err)
	}

	return product
}

func (p *ProductData) AddNewProduct(pdname string, price int, description string, amount int, uid int, bid bool, date string) error {
	var pdid int
	rows, err := p.db.Query("SELECT MAX(pd_id) FROM product")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&pdid)
		if err != nil {
			pdid = 0 // no products
		}
	}

	pdid++

	_, err = p.insert.Exec(pdid, pdname, price, description, amount, 0.0, uid, bid, date)
	return err
}

func (p *ProductData) Delete(pdid int) error {
	_, err := p._delete.Exec(pdid)
	return err
}

func (p *ProductData) UpdatePrice(pdid, price int) error {
	_, err := p.updatePrice.Exec(price, pdid)
	return err
}

func (p *ProductData) UpdateAmount(pdid, amount int) error {
	_, err := p.updateAmount.Exec(amount, pdid)
	return err
}

//func (p *ProductData) GetProductName()

func (p *ProductData) DBClose() error {
	return p.db.Close()
}
