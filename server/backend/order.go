package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"se/database"
	"time"
)

// Order is a module that handle orders
type Order struct {
	fn *database.OrderDB
}

// OrderInit return order module
func OrderInit(db *sql.DB) *Order {
	o := new(Order)
	o.fn = database.OrderDBInit(db)

	return o
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
func (o *Order) AddOrder(uid, pdid, amount int) bool {
	err := o.fn.AddOrder(uid, pdid, amount, time.Now())
	if err != nil {
		return false
	}

	return true
}

// Delete order
func (o *Order) Delete(uid, pdid int) bool {
	err := o.fn.Delete(uid, pdid)
	if err != nil {
		return false
	}

	return true
}
