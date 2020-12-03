package product

import (
	"se/database"
)

type Product struct {
	db *database.ProductData
}

func ProductInit() *Product {
	p := new(Product)
	p.db = database.ProductDataInit()
	return p
}

func (p Product) AddProduct(pdname string, price int, description string, amount int, uid int, bid bool, date string) error {
	return p.db.AddNewProduct(pdname, price, description, amount, uid, bid, date)
}

func (p *Product) ChangePrice(pdname string, price int) {
	p.db.Update(pdname)
}

func (p *Product) ChangeAmount(pdname string, amount int) {
	p.db.Update(pdname)
}

func (p *Product) ChangeDescription(pdname, description string) {
	p.db.Update(pdname)
}
