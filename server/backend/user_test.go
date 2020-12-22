package backend

// func TestUser(t *testing.T) {
// 	db := database.Open()
// 	defer db.Close()

// 	u := UserInit(db)

// 	if u.Regist("second@gmail.com", "2581473692581456", "how how") != "ok" {
// 		t.Error("cannot regist")
// 	}

// 	if _, ok := u.Login("second@gmail.com", "2581473692581456"); !ok {
// 		t.Error("regist but cannot log in.")
// 	}
// }
