package backend

import (
	"database/sql"
	"encoding/json"
	"se/database"
	"strconv"
	"strings"
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

// AddHistory add a history into user's record
func (h History) AddHistory(uid, pdid int) string {
	err := h.fn.AddHistory(uid, pdid)
	if err != nil {
		return err.Error()
	}

	return "ok"
}

// GetHistory return all historys of a user whose uid is ?
func (h History) GetHistory(uid int, amount int, newest bool) string {
	pd := h.fn.Get(uid, amount, newest)

	str, err := json.Marshal(pd)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

// Delete can delete a history user don't want to see
func (h History) Delete(uid, pid int) string {
	if err := h.fn.Delete(uid, pid); err != nil {
		return err.Error()
	}

	return "ok"
}

// DeleteSpecific delete multiple historys
func (h History) DeleteSpecific(uid int, pdid string) string {
	pdids := strings.Split(pdid, ",")

	for _, v := range pdids {
		sipd, err := strconv.Atoi(v)
		if err != nil {
			return "query contains non-integer"
		}

		if err := h.fn.Delete(uid, sipd); err != nil {
			return err.Error()
		}
	}

	return "ok"
}

// DeleteAll deletes all history of a user
func (h History) DeleteAll(uid int) string {
	if err := h.fn.DeleteAll(uid); err != nil {
		return err.Error()
	}

	return "ok"
}
