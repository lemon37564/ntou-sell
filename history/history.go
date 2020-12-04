package history

import "se/database"

type History struct {
	data      []byte
	historydb *database.HistoryDB
}

func HistoryDataInit() {
	historydb := database.HistoryDataInit()
	return
}

func (h History) GetAll(uid int) (pdid []int) { //get all history
	pdid = h.historydb.GetAll(uid)
	return pdid
}

func (h History) Delete(uid, pid int) {
	h.historydb.Delete(uid, pid)
}
