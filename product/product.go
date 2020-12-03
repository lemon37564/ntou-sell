package product

import (
	"database/sql"
	"se/database"
)

type Product struct {
	fn *database.ProductDB
}

// implemet json and logic at here

func ProductInit(db *sql.DB) *Product {
	p := new(Product)
	p.fn = database.ProductDBInit(db)
	return p
}

func (p Product) AddProduct(pdname string, price int, description string, amount int, uid int, bid bool, date string) error {
	return p.fn.AddNewProduct(pdname, price, description, amount, uid, bid, date)
}

func (p *Product) ChangePrice(pdid, price int) {
	p.fn.UpdatePrice(pdid, price)
}

func (p *Product) ChangeAmount(pdid, amount int) {
	p.fn.UpdateAmount(pdid, amount)
}

func (p *Product) ChangeDescription(pdid int, description string) {
	p.fn.UpdateDescription(pdid, description)
}
