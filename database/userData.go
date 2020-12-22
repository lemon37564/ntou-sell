package database

import (
	"database/sql"
	"log"
)

const userTable = `
CREATE TABLE IF NOT EXISTS user(
	uid int NOT NULL,
	account varchar(64) NOT NULL UNIQUE,
	password_hash varchar(64) NOT NULL,
	name varchar(64) NOT NULL,
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

type userStmt struct {
	add     *sql.Stmt
	del     *sql.Stmt
	upName  *sql.Stmt
	upPass  *sql.Stmt
	upEval  *sql.Stmt
	maxID   *sql.Stmt
	login   *sql.Stmt
	getData *sql.Stmt
	getUID  *sql.Stmt
	getAcnt *sql.Stmt
}

func userPrepare(db *sql.DB) *userStmt {
	var err error
	user := new(userStmt)

	const (
		add     = "INSERT INTO user VALUES(?,?,?,?,?);"
		del     = "DELETE FROM user WHERE uid=? AND password_hash=?;"
		upName  = "UPDATE user SET name=? WHERE uid=?"
		upPass  = "UPDATE user SET password_hash=? WHERE account=?"
		upEval  = "UPDATE user SET eval=? WHERE account=?"
		maxID   = "SELECT MAX(uid) FROM user;"
		login   = "SELECT uid FROM user WHERE account=? AND password_hash=? AND uid>0;"
		getData = "SELECT * FROM user WHERE account=? AND uid>0;"
		getUID  = "SELECT uid FROM user WHERE account=? AND uid>0;"
		getAcnt = "SELECT account FROM user WHERE uid=?;"
	)

	if user.add, err = db.Prepare(add); err != nil {
		log.Println(err)
	}

	if user.del, err = db.Prepare(del); err != nil {
		log.Println(err)
	}

	if user.upName, err = db.Prepare(upName); err != nil {
		log.Println(err)
	}

	if user.upPass, err = db.Prepare(upPass); err != nil {
		log.Println(err)
	}

	if user.upEval, err = db.Prepare(upEval); err != nil {
		log.Println(err)
	}

	if user.maxID, err = db.Prepare(maxID); err != nil {
		log.Println(err)
	}

	if user.login, err = db.Prepare(login); err != nil {
		log.Println(err)
	}

	if user.getData, err = db.Prepare(getData); err != nil {
		log.Println(err)
	}

	if user.getUID, err = db.Prepare(getUID); err != nil {
		log.Println(err)
	}

	if user.getAcnt, err = db.Prepare(getAcnt); err != nil {
		log.Println(err)
	}

	return user
}

// AddNewUser is a function for registing a new account
func (dt Data) AddNewUser(account, passwordHash, name string) error {
	var UID int

	rows, err := dt.user.maxID.Query()
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

	_, err = dt.user.add.Exec(UID, account, passwordHash, name, 0.0)
	return err
}

// DeleteUser delete data of specific user by account
func (dt Data) DeleteUser(uid int, password string) error {
	_, err := dt.user.del.Exec(uid, password)
	return err
}

// Login return user id and boolean value to check if it is valid to log in with specific account and password hash
func (dt Data) Login(account, passwordHash string) (int, bool) {
	var cnt, uid int

	rows, err := dt.user.login.Query(account, passwordHash)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&uid)
		if err != nil {
			panic(err)
		}
		cnt++
	}

	// match only one account and password_hash
	return uid, cnt == 1
}

// ChangeUserPassword updates passeword of a user by account
func (dt Data) ChangeUserPassword(account, newpass string) error {
	_, err := dt.user.upPass.Exec(account, newpass)
	return err
}

// ChangeUserName updates name of a user by account
func (dt Data) ChangeUserName(uid int, newname string) error {
	_, err := dt.user.upName.Exec(uid, newname)
	return err
}

// ChangeUserEval updates evaluation of a user by account and new eval
func (dt Data) ChangeUserEval(account string, eval float64) error {
	_, err := dt.user.upEval.Exec(account, eval)
	return err
}

// GetUIDFromAccount return user id by account
func (dt Data) GetUIDFromAccount(account string) int {
	var id int
	rows, err := dt.user.getUID.Query(account)
	if err != nil {
		log.Println(err)
		return -1
	}

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
func (dt Data) GetAccountFromUID(uid int) string {
	var account string
	rows, err := dt.user.getAcnt.Query(uid)
	if err != nil {
		log.Println(err)
		return ""
	}

	for rows.Next() {
		err = rows.Scan(&account)
		if err != nil {
			log.Println(err)
			return ""
		}
	}

	return account
}
