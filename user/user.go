package user

import (
	"database/sql"
	"se/database"
)

///// implement json in this file /////

type User struct {
	fn *database.UserDB
}

func NewUser(db *sql.DB) *User {
	u := new(User)
	u.fn = database.UserDBInit(db)

	return u
}

func (u *User) Login(account, password string) bool {
	return u.fn.Login(account, password)
}

func (u *User) Regist(account, password, name string) bool {
	err := u.fn.AddNewUser(account, password, name)
	if err != nil {
		return false // may return string here, like "account have been used"
	}

	return true
}

func (u *User) GetUserData(account string) (res string) {
	return u.fn.GetDatasFromAccount(account).String()
}

func (u *User) GetAllUserData() (res string) {
	for _, v := range u.fn.GetAllUser() {
		res += v.String()
	}

	return
}
