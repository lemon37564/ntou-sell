package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"se/database"
	"strconv"
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
func (p Product) AddProduct(sellerUID int, pdname, rawPrice, description, rawAmount, rawBid, date string) (int, error) {
	price, err := strconv.Atoi(rawPrice)
	if err != nil {
		return -1, err
	}
	if price < 1 {
		return -1, newBeErr("price cannot smaller than 1")
	}

	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return -1, err
	}
	if amount < 1 {
		return -1, newBeErr("amount cannot smaller than 1")
	}

	bid := (rawBid == "true")

	dt, err := time.Parse(TimeLayout, date)
	if err != nil {
		return -1, beError{text: "date invalid! (date format is like 2006-01-02)"}
	}

	pdid, err := p.fn.AddProduct(pdname, price, description, amount, sellerUID, bid, dt)
	if err != nil {
		if fmt.Sprint(err) == "NOT NULL constraint failed: product.seller_id" {
			return -1, beError{text: "沒有此使用者帳號!"}
		}
		return -1, err
	}
	return pdid, nil
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
func (p *Product) ChangePrice(rawPdid, rawPrice string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	price, err := strconv.Atoi(rawPrice)
	if err != nil {
		return "cannot convert " + rawPrice + " into integer", err
	}
	if price < 1 {
		return "price cannot smaller than 1", nil
	}

	err = p.fn.UpdateProductPrice(pdid, price)
	if err != nil {
		return "failed", err
	}
	return "Price has been changed", nil
}

// ChangeAmount changes amount of a specific product with it's product id
func (p *Product) ChangeAmount(rawPdid, rawAmount string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return "cannot convert " + rawAmount + " into integer", err
	}
	if amount < 1 {
		return "amount cannot smaller than 1", nil
	}

	err = p.fn.UpdateProductAmount(pdid, amount)
	if err != nil {
		return "failed", err
	}
	return "ok", nil
}

// ChangeDescription changes description of a specific product with it's product id
func (p *Product) ChangeDescription(rawPdid, description string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	err = p.fn.UpdateProductDescription(pdid, description)
	if err != nil {
		return "failed", err
	}
	return "ok", nil
}

// SetEvaluation updates eval of a specific product with it's product id
func (p *Product) SetEvaluation(rawPdid, rawEval string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	eval, err := strconv.ParseFloat(rawEval, 64)
	if err != nil {
		return "cannot convert " + rawEval + " into float", err
	}

	err = p.fn.UpdateProductEval(pdid, eval)
	if err != nil {
		return "failed", err
	}
	return "ok", nil
}

// SearchProducts return products info in json
func (p *Product) SearchProducts(name, rawMin, rawMax, rawEval string) (string, error) {
	var (
		min, max, eval int
		err            error
	)
	if rawMin == "" {
		min = 0
	} else {
		min, err = strconv.Atoi(rawMin)
		if err != nil {
			return "cannot convert " + rawMin + " into integer", err
		}
	}

	if rawMax == "" {
		max = math.MaxInt64
	} else {
		max, err = strconv.Atoi(rawMax)
		if err != nil {
			return "cannot convert " + rawMax + " into integer", err
		}
	}

	if rawEval == "" {
		eval = 0.0
	} else {
		eval, err = strconv.Atoi(rawEval)
		if err != nil {
			return "cannot convert " + rawEval + " into integer", err
		}
	}

	pds := p.fn.SearchProductWithFilter(name, min, max, eval)
	res, err := json.Marshal(pds)
	if err != nil {
		return "failed", err
	}

	return string(res), nil
}

// GetNewest return the newest product(s) in the database
func (p *Product) GetNewest(rawNumber string) (string, error) {
	number, err := strconv.Atoi(rawNumber)
	if err != nil {
		return "cannot convert " + rawNumber + " into integer", err
	}

	temp, err := json.Marshal(p.fn.NewestProduct(number))
	if err != nil {
		return "failed", err
	}
	return string(temp), nil
}

// GetProductInfo return data of a product by it's id
func (p *Product) GetProductInfo(rawPdid string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	temp, err := json.Marshal(p.fn.GetProductInfoFromPdID(pdid))
	if err != nil {
		log.Println(err)
		return "failed", err
	}

	return string(temp), nil
}

// GetSellerProduct return all product of a seller
func (p *Product) GetSellerProduct(uid int) string {
	temp, err := json.Marshal(p.fn.GetSellerProduct(uid))
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(temp)
}
