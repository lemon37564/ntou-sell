package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"se/database"
	"time"
)

type Order struct {
	fn *database.OrderDB
}

func OrderInit(db *sql.DB) *Order {
	o := new(Order)
	o.fn = database.OrderDBInit(db)

	return o
}

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

func (o *Order) AddOrder(uid, pdid, amount int) bool {
	err := o.fn.AddOrder(uid, pdid, amount, time.Now())
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
