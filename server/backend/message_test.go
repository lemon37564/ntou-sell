package backend

import (
	"testing"
)

func TestMs(t *testing.T) {
	AddMessage(1, "2", "你好")

	if res, _ := GetMessages(1, "2", "true"); res == "null" {
		t.Error("add new message but cannot found")
	}

	if res, _ := AddMessage(1, "1", "123"); res == "ok" {
		t.Error("message sent to yourself but system didn't forbid")
	}
}
