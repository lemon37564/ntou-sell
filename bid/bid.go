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
func (b Bid) Product_Bid_Current_Price(id int) int {//等等改
	return b.productDb.GetInfoFromPdID(id).Price
}

func (b Bid) GetInfoFromProductBid(pdid int) (bd database.Bid) {//等等改

	temp := b.bidDb.GetAllBid()
	return temp[pdid]
}

func (b *Bid)SetBidForBuyer(pdid,uid,money int) bool {
	if(money>b.bid)//等等改 取得競標價格
	b.bidDb.NewBidderGet(pdid,uid,money)
}

func (b Bid) Product_Bid_Evaluate(id int) float64 {
	return b.productDb.GetInfoFromPdID(id).Eval
}
func Product_Bid_User_Price() { //還沒改_

}

func (b *Bid) DeleteBid(pdid int) bool {
	
}
