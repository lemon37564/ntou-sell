package product

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"se/database"
	"strconv"
	"time"
)

const TimeLayout = "2006-01-02"

type Product struct {
	fn *database.ProductDB
}

func ProductInit(db *sql.DB) *Product {
	p := new(Product)
	p.fn = database.ProductDBInit(db)
	return p
}

//新增產品 使用: sell mod
func (p Product) AddProduct(pdname string, price int, description string, amount int, sellerUID int, bid bool, date string) (int, string) {
	dt, err := time.Parse(TimeLayout, date)
	if err != nil {
		return -1, "date invalid! (date format is like 2006-01-02)"
	}

	pdid, err := p.fn.AddNewProduct(pdname, price, description, amount, sellerUID, bid, dt)
	if err != nil {
		if fmt.Sprint(err) == "NOT NULL constraint failed: product.seller_id" {
			return -1, "沒有此使用者帳號!"
		}
		return -1, fmt.Sprint(err)
	}
	return pdid, "ok"
}

//刪除產品 使用:  mod
func (p *Product) DeleteProduct(uid int, pdname string) string {
	err := p.fn.Delete(uid, pdname)
	if err != nil {
		return fmt.Sprint(err)
	}

	return "ok"
}

//
func (p *Product) ChangePrice(pdid, price int) string {
	err := p.fn.UpdatePrice(pdid, price)
	if err != nil {
		return "Price cannot change"
	}
	return "Price has been changed"
}

func (p *Product) ChangeAmount(pdid, amount int) string {

	err := p.fn.UpdateAmount(pdid, amount)
	if err != nil {
		return "Amount cannot change"
	}
	return "Amount change success"
}

func (p *Product) ChangeDescription(pdid int, description string) string {

	err := p.fn.UpdateDescription(pdid, description)
	if err != nil {
		return "Description cannot change"
	}
	return "Description change success"
}

func (p *Product) SetEvaluation(pdid int, eval float64) string {
	err := p.fn.UpdateEval(pdid, eval)
	if err != nil {
		return "Evaluation cannot change"
	}
	return "Evaluation change success"
}

// SearchProductsByName return products info in json
func (p *Product) SearchProductsByName(name string) string {
	pds := p.fn.Search(name)

	res, err := json.Marshal(pds)
	if err != nil {
		return "Search Not Found"
	}

	return string(res)
}

func (p *Product) EnhanceSearchProductsByName(name string, minPrice, maxPrice, eval int) string {
	pds := p.fn.SearchWithFilter(name, minPrice, maxPrice, eval)

	res, err := json.Marshal(pds)
	if err != nil {
		log.Println(err)
	}

	return string(res)
}

// debugging only
func (p *Product) GetAll() string {
	pds := p.fn.All()

	res, err := json.Marshal(pds)
	if err != nil {
		log.Println(err)
	}

	return string(res)
}

func (p *Product) GetNewest(number int) string {
	temp, err := json.Marshal(p.fn.NewestProduct(number))
	if err != nil {
		log.Println(err)
		return "Fail"
	}
	return string(temp)
}

func (p *Product) GetProductInfo(uid int) string {
	//var orders string = ""
	temp, err := json.Marshal(p.fn.GetInfoFromPdID(uid))
	if err != nil {
		log.Println(err)
		return "Null"
	}
	return string(temp)

}

func (p *Product) GetProdPrice(pdid int) string { //拿價格
	return strconv.Itoa(p.fn.GetInfoFromPdID(pdid).Price)
}

func (p *Product) GetProAmount(pdid int) string { //拿數量
	return strconv.Itoa(p.fn.GetInfoFromPdID(pdid).Amount)
}

func (p *Product) GetProdDescription(pdid int) string { //拿說明
	return p.fn.GetInfoFromPdID(pdid).Description
}

func (p *Product) GetProdDate(pdid int) string { //商品釋出日期
	return p.fn.GetInfoFromPdID(pdid).Date.String()
}

func (p *Product) GetProdName(pdid int) string { //拿商品名稱
	return p.fn.GetInfoFromPdID(pdid).PdName
}

func (p *Product) GetProdEval(pdid int) string { //拿評價
	return strconv.FormatFloat(p.fn.GetInfoFromPdID(pdid).Eval, 'E', -1, 64)
}
