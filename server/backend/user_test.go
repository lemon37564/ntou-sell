package backend

import (
	"testing"
)

func TestUser(t *testing.T) {

	if Regist("second@gmail.com", "2581473692581456", "how how") != "ok" {
		t.Error("cannot regist")
	}

	if _, ok := Login("second@gmail.com", "2581473692581456"); !ok {
		t.Error("regist but cannot log in.")
	}
}

func BenchmarkLogin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Login("1234", "1234")
	}
}
