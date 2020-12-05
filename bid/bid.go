package bid

import (
	"database/sql"
	"fmt"
	"se/database"
)

type Bid struct {
	bidDb     *database.BidDB
	productDb *database.ProductDB
}

func NewBid(db *sql.DB) *Bid {
	b := new(Bid)
	b.bidDb = database.BidDataInit(db)
	return b

}

func (b Bid) ProductBidName(id int) string { //回傳商品名稱
	return b.productDb.GetInfoFromPdID(id).PdName
}
func (b Bid) ProductBidDescription(id int) string { //回傳商品描述

	return b.productDb.GetInfoFromPdID(id).Description
}
func (b Bid) ProductBidTime(id int) string { //回傳商品更新日期(非競標日期)
	return b.productDb.GetInfoFromPdID(id).Date
}

func (b Bid) ProductBidCurrentPrice(id int) int { //回傳商品目前競標價格
	return b.bidDb.GetBidByID(id).NowMoney
}

func (b Bid) GetProductBidDeadLine(pdid int) string { //回傳商品競標日期
	return b.bidDb.GetBidByID(pdid).Deadline
}
func (b *Bid) SetBidForBuyer(pdid, uid, money int) bool { //更新商品價格，目前競標者
	if money > b.bidDb.GetBidByID(pdid).NowMoney { // 取得競標價格
		b.bidDb.NewBidderGet(pdid, uid, money)
		return true
	}
	return false
}

func (b Bid) Get_Product_Bid_Evaluate(id int) float64 { //回傳評價

	return b.productDb.GetInfoFromPdID(id).Eval

}

func (b *Bid) DeleteBid(pdid int) string {
	err := b.bidDb.DeleteBid(pdid)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return "ok"
}
