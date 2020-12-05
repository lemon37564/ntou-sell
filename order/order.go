package order

import (
	"database/sql"
	"se/database"
)

type Order struct {
	fn *database.OrderDB
}

func NewOrder(db *sql.DB) *Order {
	o := new(Order)
	o.fn = database.OrderDBInit(db)

	return o
}

func (o *Order) GetOrders(uid int) (orders string) {
	//var orders string = ""

	for _, v := range o.fn.GetAllOrder(uid) {
		orders += v.String()
	}

	return
}

func (o *Order) AddOrder(uid, pdid, amount int) bool {
	err := o.fn.AddOrder(uid, pdid, amount)
	if err != nil {
		return false
	}

	return true
}

func (o *Order) Delete(uid, pdid int) bool {
	err := o.fn.Delete(uid, pdid)
	if err != nil {
		return false
	}

	return true
}
