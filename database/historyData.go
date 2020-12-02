package database

import (
	"database/sql"
)

// CREATE TABLE history(
// 	id varchar(16) NOT NULL,
// 	products varchar(2048),
// 	PRIMARY KEY(id),
// 	FOREIGN KEY(id) REFERENCES user
// );

type HistoryData struct {
	db *sql.DB

	insert  *sql.Stmt
	_delete *sql.Stmt
	update  *sql.Stmt
	_select *sql.Stmt
}

func HistoryDataInit() (*HistoryData, error) {
	history := new(HistoryData)

	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return history, err
	}
	defer db.Close()
	history.db = db

	insert, err := db.Prepare("INSERT INTO history values(?,?,?,?,?,?,?,?,?);")
	if err != nil {
		return history, err
	}
	history.insert = insert

	_delete, err := db.Prepare("DELETE FROM history where pd_id=?;")
	if err != nil {
		return history, err
	}
	history._delete = _delete

	update, err := db.Prepare("UPDATE history SET ?=?;")
	if err != nil {
		return history, err
	}
	history.update = update

	_select, err := db.Prepare("SELECT * FROM history WHERE ?=?;")
	if err != nil {
		return history, err
	}
	history._select = _select

	return history, nil
}

// wait for implementation
func (h *HistoryData) Insert() error {
	_, err := h.insert.Exec()
	return err
}

// wait for implementation
func (h *HistoryData) Delete(pdid string) error {
	_, err := h._delete.Exec(pdid)
	return err
}

// wait for implementation
func (h *HistoryData) Update(products string) error {
	return nil
}

// wait for implementation
func (h *HistoryData) Select() (string, error) {
	return "", nil
}
