package backend

import (
	"se/database"
	"testing"
)

func TestOrder(t *testing.T) {
	data := database.OpenAndInit()
	defer data.DBClose()

	uid := 0
	pdid := 0
	amount := 100000

	o := OrderInit(data)

	o.AddOrder(uid, pdid, amount)

	if res := o.GetOrders(uid); res == "null" {
		t.Error("add order but cannot found")
	}
}
