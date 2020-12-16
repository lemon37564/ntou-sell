package backend

import (
	"se/database"
	"testing"
)

func TestPd(t *testing.T) {
	db := database.Open()
	defer db.Close()

	p := ProductInit(db)

	p.AddProduct("test_product", 100, "wow", 5, 1, false, "2020-12-31")

	if res := p.SearchProductsByName("test_product"); res == "null" {
		t.Error("add new product but cannot search")
	}
}
