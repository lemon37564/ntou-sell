package account

const (
	AccountNotExist = iota
	PassWordError
	EntryValid
)

type manager struct {
	UserName, Password string
}

func (m manager) Login() int {
	if sqlName(m.UserName) {
		return AccountNotExist
	} else if sql(m.UserName) != m.Password {
		return PassWordError
	}

	return EntryValid
}

func sql(name string) string {
	return ""
}

func sqlName(name string) bool {
	return false
}
