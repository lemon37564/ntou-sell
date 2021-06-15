package backend

import (
	"testing"
)

func TestPd(t *testing.T) {

	AddProduct(1, "test_product", "100", "wow", "5", "false", "2020-12-31")

	if res, _ := SearchProducts("test_product", "0", "", ""); res == "null" {
		t.Error("add new product but cannot found")
	}
}

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchProducts("ifone", "0", "500000", "0")
	}
}
