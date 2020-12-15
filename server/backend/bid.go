package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

func (b Bid) GetProductInfo(id int) string { //回傳商品資訊

	temp, err := json.Marshal(b.productDb.GetInfoFromPdID(id))
	if err != nil {
		log.Println(err)
		return "fail to get Productinfo"
	}
	return string(temp)
}

/*func (b Bid) ProductBidDescription(id int) string { //回傳商品描述

	return b.productDb.GetInfoFromPdID(id).Description
}
func (b Bid) ProductBidTime(id int) string { //回傳商品更新日期(非競標日期)
	return b.productDb.GetInfoFromPdID(id).Date
}
*/
func (b Bid) GetProductBidInfo(id int) string { //回傳商品目前競標商品資訊
	temp, err := json.Marshal(b.bidDb.GetBidByID(id))
	if err != nil {
		log.Println(err)
		return "fail to get Bidinfo"
	}
	return string(temp)
}

/*func (b Bid) GetProductBidDeadLine(pdid int) string { //回傳商品競標日期
	return b.bidDb.GetBidByID(pdid).Deadline
}*/
func (b *Bid) SetBidForBuyer(pdid, uid, money int) bool { //更新商品價格，目前競標者
	if money > b.bidDb.GetBidByID(pdid).NowMoney { // 取得競標價格
		b.bidDb.NewBidderGet(pdid, uid, money)
		return true
	}
	return false
}

/*func (b Bid) Get_Product_Bid_Evaluate(id int) string { //回傳評價
	string := strconv.FormatFloat(b.productDb.GetInfoFromPdID(id).Eval, 'E', -1, 64)
	return string

}*/

//刪除競標
func (b *Bid) DeleteBid(pdid int) string {
	err := b.bidDb.DeleteBid(pdid)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("%v", err)
	}
	return "ok"
}
