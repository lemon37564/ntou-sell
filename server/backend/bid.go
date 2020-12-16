package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"se/database"
)

// Bid type handle bids
type Bid struct {
	fn *database.BidDB
}

// BidInit return a Bid type which handle bids
func BidInit(db *sql.DB) *Bid {
	b := new(Bid)
	b.fn = database.BidDataInit(db)
	return b
}

//GetProductInfo 回傳商品資訊
func (b Bid) GetProductInfo(pdid int) string {

	temp, err := json.Marshal(b.fn.GetAllBidProducts(pdid))
	if err != nil {
		log.Println(err)
		return "fail to get Productinfo"
	}
	return string(temp)
}

//GetProductBidInfo 回傳商品目前競標商品資訊
func (b Bid) GetProductBidInfo(pdid int) string {
	temp, err := json.Marshal(b.fn.GetBidByID(pdid))
	if err != nil {
		log.Println(err)
		return "fail to get Bidinfo"
	}
	return string(temp)
}

// SetBidForBuyer 更新商品價格，目前競標者
func (b *Bid) SetBidForBuyer(pdid, uid, money int) bool {
	if money > b.fn.GetBidByID(pdid).NowMoney { // 取得競標價格
		b.fn.NewBidderGet(pdid, uid, money)
		return true
	}
	return false
}

// DeleteBid 刪除競標
func (b *Bid) DeleteBid(pdid int) string {
	err := b.fn.DeleteBid(pdid)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("%v", err)
	}
	return "ok"
}
