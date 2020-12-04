package bid

import (
	"se/database"
	//"encoding/json"
)

type Bid struct {
	bidDb     *database.BidData
	productDb *database.ProductData
}

func BidDataInit() Bid {
	bidDb := database.BidDataInit()
	return Bid{bidDb: bidDb}

}

func (b Bid) Product_name(id int) string { //回傳商品名稱
	return b.productDb.GetProductName(id)
}
func (b Bid) Product_Description(id int) string { //回傳商品描述

	return b.productDb.GetProductDescript(id)
}
func (b Bid) Product_bid_time(id int) string {
	return b.productDb.GetBidTime(id)
}
func Product_bid_minimum(id int) int { //拿到
	return b.productDb.GetProductMinium(id)
}
func Product_Bid_Current_Price() int {
	return b.productDb.GetBidCurrentTime(id)
}

/*func Product_bid_amount() int {
	return database.Product_bid_amount()
}*/
func Product_Bid_Evaluate() {
	return database.Product_Bid_Evaluate()
}
func Product_Bid_User_Price() { //還沒改_

}
