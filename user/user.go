package user

import (
	"fmt"
	"se/database"
	"time"
)

type User struct {
	userdb *database.UserData
}

func UserInit() User {
	userdb := database.UserDataInit()
	return User{userdb: userdb}
}

func (u *User) UserListen() {
	defer u.userdb.DBClose()
	for ; ; time.Sleep(time.Second / 30) {
	}

}

func Login(account, password string) bool {
	user := database.UserDataInit()
	defer user.DBClose()

	return user.Login(account, password)
}

func Regist(account, password, name string) string {
	user := database.UserDataInit()
	defer user.DBClose()

	stat := fmt.Sprintf("%v\n", (user.AddNewUser(account, password, name)))
	fmt.Println(stat)
	return stat
}
