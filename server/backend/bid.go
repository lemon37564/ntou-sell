package backend

import (
	"encoding/json"
	"se/database"
	"strconv"
)

// GetProductBidInfo 回傳商品目前競標商品資訊
func GetProductBidInfo(pdid string) (string, error) {
	pid, err := strconv.Atoi(pdid)
	if err != nil {
		return "cannot convert " + pdid + " into integer", err
	}
	temp, err := json.Marshal(database.GetBidByID(pid))
	if err != nil {
		return "failed", err
	}
	return string(temp), nil
}

// SetBidForBuyer 更新商品價格，目前競標者
func SetBidForBuyer(uid int, pdid, money string) (string, error) {
	pid, err := strconv.Atoi(pdid)
	if err != nil {
		return "cannot convert " + pdid + " into integer", err
	}

	price, err := strconv.Atoi(money)
	if err != nil {
		return "cannot convert " + money + " into integer", err
	}

	n := database.GetProductInfoFromPdID(pid)
	if price > n.Price { // 取得競標價格
		if err := database.WonBid(uid, pid, price); err != nil {
			return "failed", err
		}

		return "ok", nil
	}

	return "price not acceptable", nil
}

// DeleteBid 刪除競標
func DeleteBid(pdid string) (string, error) {
	pid, err := strconv.Atoi(pdid)
	if err != nil {
		return "cannot convert " + pdid + " into integer", err
	}

	err = database.DeleteBid(pid)
	if err != nil {
		return "failed", err
	}
	return "ok", nil
}
