package backend

import (
	"se/database"
	"testing"
)

func TestUser(t *testing.T) {
	db := database.Open()
	defer db.Close()

	u := UserInit(db)

	u.Regist("second@gmail.com", "2581473692581456", "how how")

	if _, ok := u.Login("second@gmail.com", "2581473692581456"); !ok {
		t.Error("regist but cannot log in.")
	}
}

// func BenchmarkREG(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Regist(fmt.Sprintf("%d@gmail.com", i), "56456456", "4d5f")
// 	}
// }
