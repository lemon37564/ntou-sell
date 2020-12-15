package backend

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	Regist("second@gmail.com", "2581473692581456", "how how")
}

func BenchmarkREG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Regist(fmt.Sprintf("%d@gmail.com", i), "56456456", "4d5f")
	}
}
