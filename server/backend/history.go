package backend

import (
	"encoding/json"
	"se/database"
	"strconv"
	"strings"
)

// History contains functions to use
type History struct {
	fn *database.Data
}

// HistoryInit return handler of History
func HistoryInit(data *database.Data) *History {
	return &History{fn: data}
}

// AddHistory add a history into user's record
func (h History) AddHistory(uid int, rawPdid string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	err = h.fn.AddHistory(uid, pdid)
	if err != nil {
		return "failed", err
	}

	return "ok", nil
}

// GetHistory return all historys of a user whose uid is ?
func (h History) GetHistory(uid int, rawAmount, newest string) (string, error) {
	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return "cannot convert " + rawAmount + " into integer", err
	}

	pd := h.fn.GetAllHistory(uid, amount, newest == "true")
	str, err := json.Marshal(pd)
	if err != nil {
		return "failed", err
	}
	return string(str), nil
}

// Delete can delete a history user don't want to see
func (h History) Delete(uid int, rawPdid string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	if err = h.fn.DeleteHistory(uid, pdid); err != nil {
		return "failed", err
	}

	return "ok", nil
}

// DeleteSpecific delete multiple historys
func (h History) DeleteSpecific(uid int, pdid string) (string, error) {
	pdids := strings.Split(pdid, ",")

	for _, v := range pdids {
		sipd, err := strconv.Atoi(v)
		if err != nil {
			return "query contains non-integer", err
		}

		if err := h.fn.DeleteHistory(uid, sipd); err != nil {
			return "failed", err
		}
	}

	return "ok", nil
}

// DeleteAll deletes all history of a user
func (h History) DeleteAll(uid int) string {
	if err := h.fn.DeleteAllHistory(uid); err != nil {
		return err.Error()
	}

	return "ok"
}
