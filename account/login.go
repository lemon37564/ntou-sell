package account

import "se/database"

func Login(account, password string) bool {
	user, err := database.UserDataInit()
	if err != nil {
		panic(err)
	}
	user.Select()

	return true
}
