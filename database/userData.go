package database

import (
	"database/sql"
)

// CREATE TABLE user(
// 	id varchar(16) NOT NULL,
// 	account varchar(256) NOT NULL,
// 	password_hash varchar(64) NOT NULL,
// 	name varchar(256),
// 	eval float,
// 	PRIMARY KEY(id)
// );

type UserData struct {
	db *sql.DB

	insert  *sql.Stmt
	_delete *sql.Stmt
	update  *sql.Stmt
	_select *sql.Stmt
}

func UserDataInit() (*UserData, error) {
	user := new(UserData)

	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return user, err
	}
	defer db.Close()
	user.db = db

	insert, err := db.Prepare("INSERT INTO user values(?,?,?,?,?,?,?,?,?);")
	if err != nil {
		return user, err
	}
	user.insert = insert

	_delete, err := db.Prepare("DELETE FROM user where pd_id=?;")
	if err != nil {
		return user, err
	}
	user._delete = _delete

	update, err := db.Prepare("UPDATE user SET ?=?;")
	if err != nil {
		return user, err
	}
	user.update = update

	_select, err := db.Prepare("SELECT * FROM user WHERE ?=?;")
	if err != nil {
		return user, err
	}
	user._select = _select

	return user, nil
}

func (u *UserData) Insert(pdid string, pdname string, price int, description string, amount int, eval float64, name string, bid bool, date string) error {
	_, err := u.insert.Exec(pdid, pdname, price, description, amount, eval, name, bid, date)
	return err
}

func (u *UserData) Delete(pdid string) error {
	_, err := u._delete.Exec(pdid)
	return err
}

// wait for implementation
func (u *UserData) Update(products string) error {
	return nil
}

// wait for implementation
func (u *UserData) Select() (string, error) {
	return "", nil
}
