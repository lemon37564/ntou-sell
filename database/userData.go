package database

import (
	"database/sql"
	"fmt"
	"log"

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

type UserData struct {
	db *sql.DB

	insert     *sql.Stmt
	_delete    *sql.Stmt
	updateName *sql.Stmt
	updatePass *sql.Stmt
	updateEval *sql.Stmt
}

func UserDataInit() *UserData {
	user := new(UserData)

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	user.db = db

	user.insert, err = db.Prepare("INSERT INTO user VALUES(?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	user._delete, err = db.Prepare("DELETE FROM user WHERE account=?;")
	if err != nil {
		log.Fatal(err)
	}

	user.updateName, err = db.Prepare("UPDATE user SET name=? WHERE account=?")
	if err != nil {
		log.Fatal(err)
	}

	user.updatePass, err = db.Prepare("UPDATE user SET password_hash=? WHERE account=?")
	if err != nil {
		log.Fatal(err)
	}

	user.updateEval, err = db.Prepare("UPDATE user SET eval=? WHERE account=?")
	if err != nil {
		log.Fatal(err)
	}

	return user
}

func (u *UserData) AddNewUser(account, passwordHash, name string) error {
	var uid int
	rows, err := u.db.Query("SELECT MAX(uid) FROM user")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&uid)
		if err != nil { // no user yet
			uid = 0
		}
	}

	uid++ // for the new user

	_, err = u.insert.Exec(uid, account, passwordHash, name, 0.0)
	return err
}

func (u *UserData) DeleteUser(account string) error {
	_, err := u._delete.Exec(account)
	return err
}

// WARNING: SQL injection
func (u *UserData) Login(account, passwordHash string) bool {
	var cnt int
	rows, err := u.db.Query("SELECT * FROM user WHERE account=" + account + " AND password_hash=" + passwordHash)
	if err != nil {
		log.Fatal("logging in:", err)
	}

	for rows.Next() {
		err = rows.Scan(&cnt)
		if err != nil {
			log.Fatal("logging in:", err)
		}
	}

	// match only one account and password_hash
	return cnt == 1
}

func (u *UserData) ChangePassword(account, newpass string) error {
	_, err := u.updatePass.Exec(account, newpass)
	return err
}

func (u *UserData) ChangeName(account, newname string) error {
	_, err := u.updateName.Exec(account, newname)
	return err
}

func (u *UserData) ChangeEval(account string, eval float64) error {
	_, err := u.updateEval.Exec(account, eval)
	return err
}

func (u *UserData) DBClose() error {
	return u.db.Close()
}

type User struct {
	uid          int
	account      string
	passwordHash string
	name         string
	eval         float64
}

func (u User) String() (res string) {
	res += "user id:       " + fmt.Sprintf("%d\n", u.uid)
	res += "account:       " + u.account + "\n"
	res += "password_hash: " + u.passwordHash + "\n"
	res += "user name:     " + u.name + "\n"
	res += "evaluation:    " + fmt.Sprintf("%f\n\n", u.eval)

	return
}

// WARNING: SQL injection
func (u *UserData) GetDatasFromAccount(account string) (us User) {
	rows, err := u.db.Query("SELECT * FROM user WHERE account=" + account)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&us.uid, &us.account, &us.passwordHash, &us.name, &us.eval)
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

// WARNING: SQL injection
func (u *UserData) GetAllUser() (all []User) {
	rows, err := u.db.Query("SELECT * FROM user;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		user := *new(User)
		err = rows.Scan(&user.uid, &user.account, &user.passwordHash, &user.name, &user.eval)
		if err != nil {
			log.Fatal(err)
		}

		all = append(all, user)
	}

	return
}
