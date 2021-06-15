package backend

import (
	"testing"
)

func TestCart(t *testing.T) {

	uid := 1
	pdid := "0"
	amount := "3"

	AddProductToCart(uid, pdid, amount)
	AddProductToCart(uid, "2fff", "dsds")
	AddProductToCart(uid, pdid, "dsds")

	if GetProducts(uid) == "null" {
		t.Error("add to cart but cannot found")
	}

	RemoveProduct(uid, pdid)
	RemoveProduct(uid, "ppo")

	ModifyProductAmount(uid, pdid, amount)
	ModifyProductAmount(uid, "fdfd", "dfdf")
	ModifyProductAmount(uid, pdid, "df")
}
