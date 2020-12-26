package backend

import (
	"se/database"
	"testing"
)

func TestPd(t *testing.T) {
	data := database.OpenAndInit()
	defer data.DBClose()

	p := ProductInit(data)

	p.AddProduct("test_product", 100, "wow", 5, 1, false, "2020-12-31")

	if res := p.SearchProducts("test_product"); res == "null" {
		t.Error("add new product but cannot found")
	}
}
