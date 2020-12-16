package backend

import (
	"se/database"
	"testing"
)

func TestMs(t *testing.T) {
	db := database.Open()
	defer db.Close()

	ms := MessageInit(db)

	ms.AddMessage(1, 2, "你好")

	if ms.GetMessages(1, 2, true) == "null" {
		t.Error("add new message but cannot found")
	}

	if ms.AddMessage(1, 1, "123") == "ok" {
		t.Error("message sent to yourself but system didn't forbid")
	}
}
