package bid

import (
	"se/database"
	//"encoding/json"
)

type Bid struct {
	bidDb     *database.BidDB
	productDb *database.ProductDB
}

func BidDataInit() Bid {
	bidDb := database.BidDataInit()
	return Bid{bidDb: bidDb}

}

func (b Bid) Product_name(id int) string { //回傳商品名稱
	return b.productDb.GetInfoFromPdID(id).PdName
}
func (b Bid) Product_Description(id int) string { //回傳商品描述

	return b.productDb.GetInfoFromPdID(id).Description
}
func (b Bid) Product_bid_time(id int) string {
	return b.productDb.GetInfoFromPdID(id).Date
}

/*func Product_bid_minimum(id int) int { //拿到
	return b.productDb.GetInfoFromPdID(id).
}*/
func (b Bid) Product_Bid_Current_Price(id int) int {
	return b.productDb.GetInfoFromPdID(id).Price
}

/*func Product_bid_amount() int {
	return database.Product_bid_amount()
}*/
func (b Bid) Product_Bid_Evaluate(id int) float64 {
	return b.productDb.GetInfoFromPdID(id).Eval
}
func Product_Bid_User_Price() { //還沒改_

}
