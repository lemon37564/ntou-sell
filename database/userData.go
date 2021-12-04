package database

import (
	"database/sql"
	"log"
)

const userTable = `
CREATE TABLE IF NOT EXISTS user(
	uid int NOT NULL,
	account varchar(32) NOT NULL UNIQUE,
	password_hash varchar(32) NOT NULL,
	name varchar(32) NOT NULL,
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

var (
	userAdd     *sql.Stmt
	userDel     *sql.Stmt
	userUpName  *sql.Stmt
	userUpPass  *sql.Stmt
	userUpEval  *sql.Stmt
	userMaxID   *sql.Stmt
	userLogin   *sql.Stmt
	userGetData *sql.Stmt
	userGetUID  *sql.Stmt
	userGetAcnt *sql.Stmt
)

func userPrepare(db *sql.DB) {
	var err error

	const (
		add     = "INSERT INTO user VALUES(?,?,?,?,?);"
		del     = "DELETE FROM user WHERE uid=? AND password_hash=?;"
		upName  = "UPDATE user SET name=? WHERE uid=?"
		upPass  = "UPDATE user SET password_hash=? WHERE uid=?"
		upEval  = "UPDATE user SET eval=? WHERE account=?"
		maxID   = "SELECT MAX(uid) FROM user;"
		login   = "SELECT uid FROM user WHERE account=? AND password_hash=? AND uid>0;"
		getData = "SELECT * FROM user WHERE account=? AND uid>0;"
		getUID  = "SELECT uid FROM user WHERE account=? AND uid>0;"
		getAcnt = "SELECT account FROM user WHERE uid=?;"
	)

	if userAdd, err = db.Prepare(add); err != nil {
		panic(err)
	}

	if userDel, err = db.Prepare(del); err != nil {
		panic(err)
	}

	if userUpName, err = db.Prepare(upName); err != nil {
		panic(err)
	}

	if userUpPass, err = db.Prepare(upPass); err != nil {
		panic(err)
	}

	if userUpEval, err = db.Prepare(upEval); err != nil {
		panic(err)
	}

	if userMaxID, err = db.Prepare(maxID); err != nil {
		panic(err)
	}

	if userLogin, err = db.Prepare(login); err != nil {
		panic(err)
	}

	if userGetData, err = db.Prepare(getData); err != nil {
		panic(err)
	}

	if userGetUID, err = db.Prepare(getUID); err != nil {
		panic(err)
	}

	if userGetAcnt, err = db.Prepare(getAcnt); err != nil {
		panic(err)
	}
}

// AddNewUser is a function for registing a new account
func AddNewUser(account, passwordHash, name string) error {
	var uid int

	rows, err := userMaxID.Query()
	if err != nil {
		return err
	}

	if rows.Next() {
		err = rows.Scan(&uid)
		if err != nil {
			return err
		}
	}

	rows.Close()
	_, err = userAdd.Exec(uid+1, account, passwordHash, name, 0.0)
	return err
}

// DeleteUser delete data of specific user by account
func DeleteUser(uid int, password string) error {
	_, err := userDel.Exec(uid, password)
	return err
}

// Login return user id and boolean value to check if it is valid to log in with specific account and password hash
func Login(account, passwordHash string) (int, bool) {
	var uid int

	rows, err := userLogin.Query(account, passwordHash)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&uid)
		if err != nil {
			panic(err)
		}
		return uid, true
	}
	return -1, false
}

// ChangeUserPassword updates passeword of a user by uid
func ChangeUserPassword(uid int, newpass string) error {
	_, err := userUpPass.Exec(newpass, uid)
	return err
}

// ChangeUserName updates name of a user by account
func ChangeUserName(uid int, newname string) error {
	_, err := userUpName.Exec(newname, uid)
	return err
}

// ChangeUserEval updates evaluation of a user by account and new eval
func ChangeUserEval(account string, eval float64) error {
	_, err := userUpEval.Exec(account, eval)
	return err
}

// GetUIDFromAccount return user id by account
func GetUIDFromAccount(account string) int {
	var id int
	rows, err := userGetUID.Query(account)
	if err != nil {
		log.Println(err)
		return -1
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return -1
		}
	}

	return id
}

// GetAccountFromUID return account by user id
func GetAccountFromUID(uid int) string {
	var account string
	rows, err := userGetAcnt.Query(uid)
	if err != nil {
		log.Println(err)
		return ""
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&account)
		if err != nil {
			log.Println(err)
			return ""
		}
	}

	return account
}
