package backend

import (
	"database/sql"
	"encoding/json"
	"se/database"
)

// History contains functions to use
type History struct {
	fn *database.HistoryDB
}

// HistoryInit return handler of History
func HistoryInit(db *sql.DB) (h *History) {
	h = new(History)
	h.fn = database.HistoryDBInit(db)
	return
}

func (h History) AddHistory(uid, pdid int) string {
	err := h.fn.AddHistory(uid, pdid)
	if err != nil {
		return err.Error()
	}

	return "ok"
}

func (h History) GetHistory(uid int, amount int, newest bool) string { //get all history
	pd := h.fn.Get(uid, amount, newest)

	str, err := json.Marshal(pd)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

func (h History) Delete(uid, pid int) string {
	return h.fn.Delete(uid, pid).Error()
}
