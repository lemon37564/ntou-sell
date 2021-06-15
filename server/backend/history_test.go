package backend

import (
	"testing"
)

func TestHistory(t *testing.T) {
	AddHistory(2, "2")

	if res, _ := GetHistory(2, "2", "true"); res == "null" {
		t.Error("add history but cannot found")
	}
}
