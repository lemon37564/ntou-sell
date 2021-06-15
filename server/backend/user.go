package backend

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"se/database"
	"unicode"
)

// Login return user id and is it valid to login
func Login(account, password string) (int, bool) {
	hash := sha256Hash(password)
	return database.Login(account, hash)
}

// Regist let user regist his own account
func Regist(account, password, name string) string {
	if containCh(account) {
		return "帳號不能含有中文"
	}

	if containCh(password) {
		return "密碼不能為中文"
	}

	hash := sha256Hash(password)

	err := database.AddNewUser(account, hash, name)
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
func DeleteUser(uid int, password string) string {
	err := database.DeleteUser(uid, password)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return "ok"
}

// ChangePassword changes user's password
func ChangePassword(uid int, oldPassword, newPassword string) string {

	account := database.GetAccountFromUID(uid)

	_, ok := database.Login(account, oldPassword)
	if !ok {
		return "舊密碼錯誤"
	}

	hash := sha256Hash(newPassword)
	err := database.ChangeUserPassword(uid, hash)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return "ok"
}

// ChangeName changes user name
func ChangeName(uid int, newname string) string {
	err := database.ChangeUserName(uid, newname)
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
