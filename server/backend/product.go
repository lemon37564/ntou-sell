package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"se/database"
	"time"
)

// Product is a module that handle products
type Product struct {
	fn *database.Data
}

// ProductInit return product module
func ProductInit(data *database.Data) *Product {
	return &Product{fn: data}
}

// AddProduct adds a product with multiple parameters
func (p Product) AddProduct(pdname string, price int, description string, amount int, sellerUID int, bid bool, date string) (int, string) {
	dt, err := time.Parse(TimeLayout, date)
	if err != nil {
		return -1, "date invalid! (date format is like 2006-01-02)"
	}

	pdid, err := p.fn.AddProduct(pdname, price, description, amount, sellerUID, bid, dt)
	if err != nil {
		if fmt.Sprint(err) == "NOT NULL constraint failed: product.seller_id" {
			return -1, "沒有此使用者帳號!"
		}
		return -1, fmt.Sprint(err)
	}
	return pdid, "ok"
}

// DeleteProduct deletes a product with seller_uid and product name
// This may me cause some problem, need to fix
func (p *Product) DeleteProduct(uid int, pdname string) string {
	err := p.fn.DeleteProduct(uid, pdname)
	if err != nil {
		return fmt.Sprint(err)
	}

	return "ok"
}

// ChangePrice changes price of a specific product with it's product id
func (p *Product) ChangePrice(pdid, price int) string {
	err := p.fn.UpdateProductPrice(pdid, price)
	if err != nil {
		return "Price cannot change"
	}
	return "Price has been changed"
}

// ChangeAmount changes amount of a specific product with it's product id
func (p *Product) ChangeAmount(pdid, amount int) string {

	err := p.fn.UpdateProductAmount(pdid, amount)
	if err != nil {
		return "Amount cannot change"
	}
	return "Amount change success"
}

// ChangeDescription changes description of a specific product with it's product id
func (p *Product) ChangeDescription(pdid int, description string) string {

	err := p.fn.UpdateProductDescription(pdid, description)
	if err != nil {
		return "Description cannot change"
	}
	return "Description change success"
}

// SetEvaluation updates eval of a specific product with it's product id
func (p *Product) SetEvaluation(pdid int, eval float64) string {
	err := p.fn.UpdateProductEval(pdid, eval)
	if err != nil {
		return "Evaluation cannot change"
	}
	return "Evaluation change success"
}

// SearchProductsByName return products info in json
func (p *Product) SearchProductsByName(name string) string {
	pds := p.fn.SearchProduct(name)

	res, err := json.Marshal(pds)
	if err != nil {
		return err.Error()
	}

	return string(res)
}

// EnhanceSearchProductsByName is a advanced function of normal search function
// it can limit the maximum price, minimum price and evaluation
func (p *Product) EnhanceSearchProductsByName(name string, minPrice, maxPrice, eval int) string {
	pds := p.fn.SearchProductWithFilter(name, minPrice, maxPrice, eval)

	res, err := json.Marshal(pds)
	if err != nil {
		log.Println(err)
	}

	return string(res)
}

// GetNewest return the newest product(s) in the database
func (p *Product) GetNewest(number int) string {
	temp, err := json.Marshal(p.fn.NewestProduct(number))
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	return string(temp)
}

// GetProductInfo return data of a product by it's id
func (p *Product) GetProductInfo(pdid int) string {
	//var orders string = ""
	temp, err := json.Marshal(p.fn.GetProductInfoFromPdID(pdid))
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(temp)
}
