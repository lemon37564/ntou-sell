package backend

import (
	"se/database"
	"testing"
)

func TestCart(t *testing.T) {
	db := database.Open()
	defer db.Close()

	c := CartInit(db)

	uid := 2
	pdid := 1
	amount := 50

	c.AddProductToCart(uid, pdid, amount)

	if res := c.GetProducts(uid); res == "null" {
		t.Error("add to cart but cannot found")
	}
}
