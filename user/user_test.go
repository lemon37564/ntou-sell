package user

import (
	"testing"
)

func TestRegister(t *testing.T) {
	b := Regist("second@gmail.com", "2581473692581456", "how how")
	if b != "" {
		panic(b)
	}
}
