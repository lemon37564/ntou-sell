package backend

import (
	"se/database"
	"testing"
)

func TestPd(t *testing.T) {
	data := database.OpenAndInit()
	defer data.DBClose()

	p := ProductInit(data)

	p.AddProduct(1, "test_product", "100", "wow", "5", "false", "2020-12-31")

	if res, _ := p.SearchProducts("test_product", "0", "", ""); res == "null" {
		t.Error("add new product but cannot found")
	}
}
