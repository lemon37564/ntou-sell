package user

import (
	"database/sql"
	"fmt"
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

func (u *User) Regist(account, password, name string) string {
	err := u.fn.AddNewUser(account, password, name)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	return "ok"
}

func (u *User) DeleteUser(account, password string) string {
	err := u.fn.DeleteUser(account, password)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	return "ok"
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
