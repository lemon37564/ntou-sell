package backend

import (
	"se/database"
	"testing"
	"time"
)

func TestCart(t *testing.T) {
	db := database.Open()
	defer db.Close()

	c := CartInit(db)

	uid := 1
	pdid := 2
	amount := 3

	c.AddProductToCart(uid, pdid, amount)

	time.Sleep(time.Second)

	t.Log(c.Debug())

	if c.GetProducts(uid) == "null" {
		t.Error("add to cart but cannot found")
	}

	// if c.TotalCount(uid) == "0" {
	// 	t.Error("add to cart but cannot found")
	// }

}
