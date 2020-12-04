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

func (s *Sell) SetProductpdid(pdid int, pdname string, price int, description string, amount int, sellerID int, bid bool, date string) bool {
	pid, err := s.fn.AddNewProduct(pdname, price, description, amount, sellerID, bid, date)
	if err != nil {
		return false
	}

	if bid { //等傳
		s.fn2.AddNewBid(s.fn.GetInfoFromPdID(pid), date, inimoney, sellerID)
	}
	return true
}
