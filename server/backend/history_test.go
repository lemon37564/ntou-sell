package backend

import (
	"se/database"
	"testing"
)

func TestHistory(t *testing.T) {
	data := database.OpenAndInit()
	defer data.DBClose()

	h := HistoryInit(data)

	h.AddHistory(2, 2)

	if res := h.GetHistory(2, 2, true); res == "null" {
		t.Error("add history but cannot found")
	}
}
