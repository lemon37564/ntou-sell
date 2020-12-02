package account

import "se/database"

func regist(account, password string) bool {
	user, err := database.UserDataInit()
	if err != nil {
		panic(err)
	}
	user.Insert()
	return false
}
