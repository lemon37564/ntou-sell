package backend

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"se/database"
	"unicode"
)

// User is a module that handle users
type User struct {
	fn *database.Data
}

// UserInit return a user module
func UserInit(data *database.Data) *User {
	return &User{fn: data}
}

// Login return user id and is it valid to login
func (u *User) Login(account, password string) (int, bool) {
	hash := sha256Hash(password)
	return u.fn.Login(account, hash)
}

// Regist let user regist his own account
func (u *User) Regist(account, password, name string) string {
	if containCh(account) {
		return "帳號不能含有中文"
	}

	if containCh(password) {
		return "密碼不能為中文"
	}

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

	err := u.fn.ChangeUserPassword(account, newPassword)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return "ok"
}

// ChangeName changes user name
func (u *User) ChangeName(uid int, newname string) string {
	err := u.fn.ChangeUserName(uid, newname)
	if err != nil {
		log.Println(err)
		return "failed"
	}

	return "ok"
}

const salt = "ntou-sell"

func sha256Hash(key string) string {
	key += salt

	hasher := sha256.New()
	hasher.Write([]byte(key))

	t := hasher.Sum(nil)
	return base64.URLEncoding.EncodeToString(t)
}

func containCh(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}

	return false
}
