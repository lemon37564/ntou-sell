package backend

import (
	"se/database"
	"testing"
	"time"
)

func TestCart(t *testing.T) {
	data := database.OpenAndInit()
	defer data.DBClose()

	c := CartInit(data)

	uid := 1
	pdid := "0"
	amount := "3"

	c.AddProductToCart(uid, pdid, amount)

	time.Sleep(time.Second)

	if c.GetProducts(uid) == "null" {
		t.Error("add to cart but cannot found")
	}

}
