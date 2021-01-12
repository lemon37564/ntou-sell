package backend

import (
	"se/database"
	"testing"
)

func TestMs(t *testing.T) {
	data := database.OpenAndInit()
	defer data.DBClose()

	ms := MessageInit(data)

	ms.AddMessage(1, "2", "你好")

	if res, _ := ms.GetMessages(1, "2", "true"); res == "null" {
		t.Error("add new message but cannot found")
	}

	if res, _ := ms.AddMessage(1, "1", "123"); res == "ok" {
		t.Error("message sent to yourself but system didn't forbid")
	}
}
