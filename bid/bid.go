package bid

import (
	"se/database"
	//"encoding/json"
)

type bid struct {
	db database.BidData
}

func (b *bid) Product_name(id int) string { //回傳商品名稱
	return database.ProductName(id)
}
func Product_Description(id int) string { //回傳商品描述
	return database.Product_Description()
}
func Product_bid_time(id int) string {
	return database.Product_bid_time()
}
func Product_bid_minimum(id int) int {
	return database.Product_Bid_Minimum()
}
func Product_Bid_Current_Price() int {
	return database.Product_Bid_Current_Price()
}
func Product_bid_amount() int {
	return database.Product_bid_amount()
}
func Product_Bid_Evaluate() {
	return database.Product_Bid_Evaluate()
}
func Product_Bid_User_Price() {
	return database.Product_Bid_User_Price()
}
