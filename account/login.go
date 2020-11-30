package account

import "se/database"

const (
	AccountNotExist = iota
	PassWordError
	EntryValid
)

type manager struct {
	UserName, Password string
}

func (m manager) Login() int {
	if database.SqlName(m.UserName) {
		return AccountNotExist
	} else if database.Sql(m.UserName) != m.Password {
		return PassWordError
	}

	return EntryValid
}
