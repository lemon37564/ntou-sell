package backend

import (
	"se/database"
	"testing"
)

func TestOrder(t *testing.T) {
	db := database.Open()
	defer db.Close()

	uid := 0
	pdid := 0
	amount := 100000

	o := OrderInit(db)

	o.AddOrder(uid, pdid, amount)

	if res := o.GetOrders(uid); res == "null" {
		t.Error("add order but cannot found")
	}
}
