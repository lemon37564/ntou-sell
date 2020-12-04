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
	o.fn = database.OrderDBInit()
}
func (o *Order) GetOrders(uid int) (ods []database.Order) (orders string) {
	
	return o.fn.GetAllOrder(uid)
}

func (o *Order) AddOrder(uid,pdid,amount int) bool {
	err := o.fn.AddOrder(uid,pdid,amount)
	if err != nil {
		return false
	}

	return true
}

func (o *Order) Delete (uid,pdid int) bool {
	err := o.fn.Delete(uid,pdid) 
	if err != nil {
		return false
	}

	return true
}
