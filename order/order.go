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
func (o *Order) GetOrder(uid) []byte {
	return o.data
}
