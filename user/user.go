package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"se/database"
)

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
		log.Println(err)
		if fmt.Sprint(err) == "UNIQUE constraint failed: user.account" {
			return "此帳號已有被註冊過!"
		}
		return fmt.Sprint(err)
	}

	return "ok"
}

func (u *User) DeleteUser(account, password string) string {
	err := u.fn.DeleteUser(account, password)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("%v", err)
	}

	return "ok"
}

func (u *User) GetUserData(account string) string {
	res, _ := json.Marshal(u.fn.GetDatasFromAccount(account))
	return string(res)
}

func (u *User) GetUIDFromAccount(account string) int {
	return u.fn.GetUIDFromAccount(account)
}

func (u *User) GetAllUserData() string {
	res, _ := json.Marshal(u.fn.GetAllUser())
	return string(res)
}
