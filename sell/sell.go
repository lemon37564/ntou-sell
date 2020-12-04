package sell

import (
	"database/sql"
	"se/database"
)

type Sell struct {
	fn  *database.ProductDB
	fn2 *database.BidDB
}

func NewProduct(db *sql.DB) (s *Sell) {
	s.fn = database.ProductDBInit(db)
	return
}

func (s *Sell) SetProductpdid(pdname string, price int, description string, amount int, sellerID int, bid bool, date string, dateLine string) string { //當在競標時為競標價格
	pid, err := s.fn.AddNewProduct(pdname, price, description, amount, sellerID, bid, date)
	if err != nil {
		return "Something Wrong when you enter product info"
	}

	if bid { //等傳

		err := s.fn2.AddNewBid(pid, dateLine, price, sellerID)
		if err != nil {
			return "Something Wrong in bid info"
		}
	}

	return "Success Add Product"
}
