package backend

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"log"
	"se/database"
)

// User is a module that handle users
type User struct {
	fn *database.UserDB
}

// UserInit return a user module
func UserInit(db *sql.DB) *User {
	u := new(User)
	u.fn = database.UserDBInit(db)

	return u
}

// Login return user id and is it valid to login
func (u *User) Login(account, password string) (int, bool) {
	hash := sha256Hash(password)
	return u.fn.Login(account, hash)
}

// Regist let user regist his own account
func (u *User) Regist(account, password, name string) string {
	hash := sha256Hash(password)

	err := u.fn.AddNewUser(account, hash, name)
	if err != nil {
		log.Println(err)
		if err.Error() == "UNIQUE constraint failed: user.account" {
			return "此帳號已有被註冊過!"
		}
		return err.Error()
	}

	return "ok"
}

// DeleteUser simple delete his account
func (u *User) DeleteUser(uid int, password string) string {
	err := u.fn.DeleteUser(uid, password)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return "ok"
}

// ChangePassword changes user's password
func (u *User) ChangePassword(uid int, oldPassword, newPassword string) string {

	account := u.fn.GetAccountFromUID(uid)

	_, ok := u.Login(account, oldPassword)
	if !ok {
		return "舊密碼錯誤"
	}

	err := u.fn.ChangePassword(account, newPassword)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return "ok"
}

// ChangeName changes user name
func (u *User) ChangeName(uid int, newname string) string {
	err := u.fn.ChangeName(uid, newname)
	if err != nil {
		log.Println(err)
		return "failed"
	}

	return "ok"
}

// GetUserData return user data with account
func (u *User) GetUserData(account string) string {
	res, _ := json.Marshal(u.fn.GetDatasFromAccount(account))
	return string(res)
}

// GetAllUserData return all users (debugging only)
func (u *User) GetAllUserData() string {
	res, _ := json.Marshal(u.fn.GetAllUser())
	return string(res)
}

func sha256Hash(key string) string {
	salt := "ntou-sell"
	key += salt

	hasher := sha256.New()
	hasher.Write([]byte(key))

	t := hasher.Sum(nil)
	return base64.URLEncoding.EncodeToString(t)
}
