package product

import (
	"se/database"
)

type Product struct {
	db database.ProductData
}

func (p *Product) AddProduct(price, amount int, pdname, desciption string, bid bool) error {
	err := p.db.Insert("", pdname, price, desciption, amount, 0, "", bid, "")
	return err
}

func (p Product) Price(pdname string) int {
	p.db.Select()
	return 0
}

func (p Product) Amount() int {
	p.db.Select()
	return 0
}

func (p Product) Description() string {
	p.db.Select()
	return ""
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
