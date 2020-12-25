package backend

import (
	"se/database"
	"testing"
)

func TestUser(t *testing.T) {
	data := database.OpenAndInit()
	defer data.DBClose()

	u := UserInit(data)

	if u.Regist("second@gmail.com", "2581473692581456", "how how") != "ok" {
		t.Error("cannot regist")
	}

	if _, ok := u.Login("second@gmail.com", "2581473692581456"); !ok {
		t.Error("regist but cannot log in.")
	}
}
