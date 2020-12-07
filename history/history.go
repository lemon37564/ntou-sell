package history

import (
	"database/sql"
	"se/database"
)

type History struct {
	historydb *database.HistoryDB
}

func NewHistory(db *sql.DB) (u *History) {
	u = new(History)
	u.historydb = database.HistoryDBInit(db)
	return
}

func (h History) GetAllHistory(uid int) (pdid []int) { //get all history
	pdid = h.historydb.GetAll(uid)
	return pdid
}

func (h History) Delete(uid, pid int) {
	h.historydb.Delete(uid, pid)
}
