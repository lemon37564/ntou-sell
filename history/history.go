package history

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"se/database"
)

type History struct {
	historydb *database.HistoryDB
	productdb *database.ProductDB
}

func NewHistory(db *sql.DB) (u *History) {
	u = new(History)
	u.historydb = database.HistoryDBInit(db)
	return
}

func (h History) AddHistory(uid, pdid int) string {
	err := h.historydb.AddHistory(uid, pdid)
	if err != nil {
		return fmt.Sprint(err)
	}

	return "ok"
}

func (h History) GetHistory(uid int, amount int) string { //get all history
	var temp []database.Product

	pdid := h.historydb.Get(uid, amount)

	for _, v := range pdid {
		temp = append(temp, h.productdb.GetInfoFromPdID(v))
	}

	str, err := json.Marshal(temp)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func (h History) GetAll() string {
	all := h.historydb.GetAll()
	res, err := json.Marshal(all)
	if err != nil {
		panic(err)
	}

	return string(res)
}

func (h History) Delete(uid, pid int) {
	h.historydb.Delete(uid, pid)
}
