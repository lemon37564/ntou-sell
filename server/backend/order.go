package backend

import (
	"encoding/json"
	"log"
	"se/database"
	"strconv"
	"time"
)

// GetOrders Return orders of a specific user
func GetOrders(uid int) string {
	//var orders string = ""
	pds := database.GetAllOrder(uid)
	res, err := json.Marshal(pds)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return string(res)
}

// AddOrder adds a order of a specific user with product id and amount
func AddOrder(uid int, rawPdid, rawAmount string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return "cannot convert " + rawAmount + " into integer", err
	}

	err = database.AddOrder(uid, pdid, amount, time.Now())
	if err != nil {
		return "false", err
	}

	return "true", nil
}

// Delete order
func DeleteOrder(uid int, rawPdid string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	err = database.DeleteOrder(uid, pdid)
	if err != nil {
		return "false", err
	}

	return "true", nil
}
