package user

import (
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

func (u *User) Login(account, password string) bool {
	return u.userdb.Login(account, password)
}

func (u *User) Regist(account, password, name string) error {
	return u.userdb.AddNewUser(account, password, name)
}

func (u *User) GetUserData(account string) (res string) {
	return u.userdb.GetDatasFromAccount(account).String()
}

func (u *User) GetAllUserData() (res string) {
	for _, v := range u.userdb.GetAllUser() {
		res += v.String()
	}

	return
}
