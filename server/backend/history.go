package backend

import (
	"encoding/json"
	"se/database"
	"strconv"
	"strings"
)

// AddHistory add a history into user's record
func AddHistory(uid int, rawPdid string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	err = database.AddHistory(uid, pdid)
	if err != nil {
		return "failed", err
	}

	return "ok", nil
}

// GetHistory return all historys of a user whose uid is ?
func GetHistory(uid int, rawAmount, newest string) (string, error) {
	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return "cannot convert " + rawAmount + " into integer", err
	}

	pd := database.GetAllHistory(uid, amount, newest == "true")
	str, err := json.Marshal(pd)
	if err != nil {
		return "failed", err
	}
	return string(str), nil
}

// Delete can delete a history user don't want to see
func DeleteHistory(uid int, rawPdid string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	if err = database.DeleteHistory(uid, pdid); err != nil {
		return "failed", err
	}

	return "ok", nil
}

// DeleteSpecific delete multiple historys
func DeleteSpecificHistory(uid int, pdid string) (string, error) {
	pdids := strings.Split(pdid, ",")

	for _, v := range pdids {
		sipd, err := strconv.Atoi(v)
		if err != nil {
			return "query contains non-integer", err
		}

		if err := database.DeleteHistory(uid, sipd); err != nil {
			return "failed", err
		}
	}

	return "ok", nil
}

// DeleteAll deletes all history of a user
func DeleteAllHistory(uid int) string {
	if err := database.DeleteAllHistory(uid); err != nil {
		return err.Error()
	}

	return "ok"
}
