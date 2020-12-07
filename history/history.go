package history

import (
	"database/sql"
	"encoding/json"
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

func (h History) GetAllHistory(uid int) string { //get all history
	var temp []database.Product

	pdid := h.historydb.GetAll(uid)

	for _, v := range pdid {

		temp = append(temp, h.productdb.GetInfoFromPdID(v))
	}
	str, err := json.Marshal(temp)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func (h History) Delete(uid, pid int) {
	h.historydb.Delete(uid, pid)
}
