package sell

import (
	"database/sql"
	"log"
	"se/database"
)

type Sell struct {
	fn  *database.ProductDB
	fn2 *database.BidDB
}

func NewSell(db *sql.DB) (s *Sell) {
	s = new(Sell)
	s.fn = database.ProductDBInit(db)
	return
}

func (s *Sell) SetProductpdid(pdname string, price int, description string, amount int, account string, sellerID int, bid bool, date string, dateLine string) string { //當在競標時為競標價格
	pid, err := s.fn.AddNewProduct(pdname, price, description, amount, account, bid, date)
	if err != nil {
		log.Println(err)
		return "Something Wrong when you enter product info"
	}

	if bid { //等傳

		err := s.fn2.AddNewBid(pid, dateLine, price, sellerID)
		if err != nil {
			log.Println(err)
			return "Something Wrong in bid info"
		}
	}

	return "Success Add Product"
}
