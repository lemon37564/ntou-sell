package backend

import (
	"encoding/json"
	"log"
	"se/database"
	"strconv"
	"time"
)

// Order is a module that handle orders
type Order struct {
	fn *database.Data
}

// OrderInit return order module
func OrderInit(data *database.Data) *Order {
	return &Order{fn: data}
}

// GetOrders Return orders of a specific user
func (o *Order) GetOrders(uid int) string {
	//var orders string = ""
	pds := o.fn.GetAllOrder(uid)
	res, err := json.Marshal(pds)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return string(res)
}

// AddOrder adds a order of a specific user with product id and amount
func (o *Order) AddOrder(uid int, rawPdid, rawAmount string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return "cannot convert " + rawAmount + " into integer", err
	}

	err = o.fn.AddOrder(uid, pdid, amount, time.Now())
	if err != nil {
		return "false", err
	}

	return "true", nil
}

// Delete order
func (o *Order) Delete(uid int, rawPdid string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	err = o.fn.DeleteOrder(uid, pdid)
	if err != nil {
		return "false", err
	}

	return "true", nil
}
