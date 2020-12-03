package database

import (
	"database/sql"
	"log"
)

const userTable = `CREATE TABLE user(
	id varchar(16) NOT NULL,
	account varchar(256) NOT NULL,
	password_hash varchar(64) NOT NULL,
	name varchar(256),
	eval float,
	PRIMARY KEY(id)
);`

type UserData struct {
	db *sql.DB

	insert  *sql.Stmt
	_delete *sql.Stmt
	update  *sql.Stmt
}

func UserDataInit() *UserData {
	user := new(UserData)

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	user.db = db

	insert, err := db.Prepare("INSERT INTO user values(?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	user.insert = insert

	_delete, err := db.Prepare("DELETE FROM user where id=?;")
	if err != nil {
		log.Fatal(err)
	}
	user._delete = _delete

	return user
}

func (u *UserData) AddNewUser(account, passwordHash, name string) error {
	if len(passwordHash) != 16 {
		return HashValError{length: len(passwordHash)}
	}

	id := "temp"

	_, err := u.insert.Exec(id, account, name, passwordHash, 0.0)
	return err
}

func (u *UserData) DeleteUser(id string) error {
	_, err := u._delete.Exec(id)
	return err
}

// WARNING: SQL injection
func (u *UserData) Login(account, passwordHash string) bool {
	var cnt int
	rows, err := u.db.Query("select * from user where account=" + account + " and password_hash=" + passwordHash)
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

// wait for implementation
func (u *UserData) Update(products string) error {
	return nil
}

func (u *UserData) DBClose() error {
	return u.db.Close()
}
