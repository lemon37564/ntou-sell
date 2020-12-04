package history

import "se/database"

type History struct {
	data      []byte
	historydb *database.HistoryData
}

func HistoryDataInit() {
	historydb := database.HistoryDataInit()
	return
}

func (h History) GetHistory() []byte {
	return h.data
}

func (h History) Delete(uid, pid int) {

}
