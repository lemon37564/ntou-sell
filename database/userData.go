package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const userTable = `CREATE TABLE user(
						uid int NOT NULL,
						account varchar(256) NOT NULL UNIQUE,
						password_hash varchar(64) NOT NULL,
						name varchar(256) NOT NULL,
						eval float,
						PRIMARY KEY(uid)
					);`

// User type contains data of single user
type User struct {
	UID          int
	Account      string
	PasswordHash string
	Name         string
	Eval         float64
}

func (u User) String() (res string) {
	res += "user id:       " + fmt.Sprintf("%d\n", u.UID)
	res += "account:       " + u.Account + "\n"
	res += "password hash: " + u.PasswordHash + "\n"
	res += "user name:     " + u.Name + "\n"
	res += "evaluation:    " + fmt.Sprintf("%f\n\n", u.Eval)

	return
}

// UserDB contain functions to use
type UserDB struct {
	insert     *sql.Stmt
	_delete    *sql.Stmt
	updateName *sql.Stmt
	updatePass *sql.Stmt
	updateEval *sql.Stmt
	maxID      *sql.Stmt
	login      *sql.Stmt
	getData    *sql.Stmt
	allUser    *sql.Stmt
}

// UserDBInit initialize all functions in user database
func UserDBInit(db *sql.DB) *UserDB {
	var err error
	user := new(UserDB)

	user.insert, err = db.Prepare("INSERT INTO user VALUES(?,?,?,?,?);")
	if err != nil {
		panic(err)
	}

	user._delete, err = db.Prepare("DELETE FROM user WHERE account=?;")
	if err != nil {
		panic(err)
	}

	user.updateName, err = db.Prepare("UPDATE user SET name=? WHERE account=?")
	if err != nil {
		panic(err)
	}

	user.updatePass, err = db.Prepare("UPDATE user SET password_hash=? WHERE account=?")
	if err != nil {
		panic(err)
	}

	user.updateEval, err = db.Prepare("UPDATE user SET eval=? WHERE account=?")
	if err != nil {
		panic(err)
	}

	user.maxID, err = db.Prepare("SELECT MAX(UID) FROM user")
	if err != nil {
		panic(err)
	}

	user.login, err = db.Prepare("SELECT COUNT(*) FROM user WHERE account=? AND password_hash=?;")
	if err != nil {
		panic(err)
	}

	user.getData, err = db.Prepare("SELECT * FROM USER WHERE account=?;")
	if err != nil {
		panic(err)
	}

	user.allUser, err = db.Prepare("SELECT * FROM USER WHERE uid>0;")
	if err != nil {
		panic(err)
	}

	return user
}

// AddNewUser is a function for registing a new account
func (u *UserDB) AddNewUser(account, passwordHash, name string) error {

	var UID int

	rows, err := u.maxID.Query()
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&UID)
		if err != nil {
			return err
		}
	}

	UID++

	_, err = u.insert.Exec(UID, account, passwordHash, name, 0.0)
	return err
}

// DeleteUser delete data of specific user by account
func (u *UserDB) DeleteUser(account string) error {
	_, err := u._delete.Exec(account)
	return err
}

// Login return boolean value to check if it is valid to log in with specific account and password hash
func (u *UserDB) Login(account, passwordHash string) bool {
	var cnt int

	rows, err := u.login.Query(account, passwordHash)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&cnt)
		if err != nil {
			panic(err)
		}
	}

	// match only one account and password_hash
	return cnt == 1
}

// ChangePassword updates passeword of a user by account
func (u *UserDB) ChangePassword(account, newpass string) error {
	_, err := u.updatePass.Exec(account, newpass)
	return err
}

// ChangeName updates name of a user by account
func (u *UserDB) ChangeName(account, newname string) error {
	_, err := u.updateName.Exec(account, newname)
	return err
}

// ChangeEval updates evaluation of a user by account and new eval
func (u *UserDB) ChangeEval(account string, eval float64) error {
	_, err := u.updateEval.Exec(account, eval)
	return err
}

// GetDatasFromAccount return data of user, matching by account
// it can also use for getting id of user from account
func (u *UserDB) GetDatasFromAccount(account string) (us User) {
	rows, err := u.getData.Query(account)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&us.UID, &us.Account, &us.PasswordHash, &us.Name, &us.Eval)
		if err != nil {
			panic(err)
		}
	}

	return
}

// GetAllUser return data of all user in the database
func (u *UserDB) GetAllUser() (all []User) {
	rows, err := u.allUser.Query()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		user := *new(User)
		err = rows.Scan(&user.UID, &user.Account, &user.PasswordHash, &user.Name, &user.Eval)
		if err != nil {
			panic(err)
		}

		all = append(all, user)
	}

	return
}
