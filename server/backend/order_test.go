package backend

import (
	"testing"
)

func TestOrder(t *testing.T) {

	uid := 0
	pdid := "0"
	amount := "100000"

	AddOrder(uid, pdid, amount)

	if res := GetOrders(uid); res == "null" {
		t.Error("add order but cannot found")
	}
}
