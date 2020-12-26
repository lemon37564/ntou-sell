package backend

import (
	"encoding/json"
	"log"
	"se/database"
	"strconv"
)

// Bid type handle bids
type Bid struct {
	fn *database.Data
}

// BidInit return a Bid type which handle bids
func BidInit(data *database.Data) *Bid {
	return &Bid{fn: data}
}

//GetProductBidInfo 回傳商品目前競標商品資訊
func (b Bid) GetProductBidInfo(pdid string) (string, error) {
	pid, err := strconv.Atoi(pdid)
	if err != nil {
		return "cannot convert " + pdid + " into integer", err
	}
	temp, err := json.Marshal(b.fn.GetBidByID(pid))
	if err != nil {
		log.Println(err)
		return "fail to get Bidinfo", nil
	}
	return string(temp), nil
}

// SetBidForBuyer 更新商品價格，目前競標者
func (b *Bid) SetBidForBuyer(uid int, pdid, money string) (string, error) {
	pid, err := strconv.Atoi(pdid)
	if err != nil {
		return "cannot convert " + pdid + " into integer", err
	}

	price, err := strconv.Atoi(money)
	if err != nil {
		return "cannot convert " + money + " into integer", err
	}

	if price > b.fn.GetBidByID(pid).NowMoney { // 取得競標價格
		if err := b.fn.WonBid(pid, uid, price); err != nil {
			log.Println(err)
			return "WonBid failed", err
		}

		return "ok", nil
	}

	return "price not acceptable", nil
}

// DeleteBid 刪除競標
func (b *Bid) DeleteBid(pdid string) (string, error) {
	pid, err := strconv.Atoi(pdid)
	if err != nil {
		return "cannot convert " + pdid + " into integer", err
	}

	err = b.fn.DeleteBid(pid)
	if err != nil {
		log.Println(err)
		return "failed", err
	}
	return "ok", nil
}
