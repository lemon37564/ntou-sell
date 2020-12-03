package database

import (
	"database/sql"
	"log"
)

const historyTable = `CREATE TABLE history(
						uid int NOT NULL,
						products varchar(2048),
						PRIMARY KEY(uid),
						FOREIGN KEY(uid) REFERENCES user
					);`

type HistoryData struct {
	db *sql.DB

	insert  *sql.Stmt
	_delete *sql.Stmt
	update  *sql.Stmt
	_select *sql.Stmt
}

func HistoryDataInit() *HistoryData {
	history := new(HistoryData)

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	history.db = db

	history.insert, err = db.Prepare("INSERT INTO history values(?,?);")
	if err != nil {
		log.Fatal(err)
	}

	history._delete, err = db.Prepare("DELETE FROM history where pd_id=?;")
	if err != nil {
		log.Fatal(err)
	}

	history.update, err = db.Prepare("UPDATE history SET ?=?;")
	if err != nil {
		log.Fatal(err)
	}

	return history
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

// always use this function at the end
func (h *HistoryData) DBClose() error {
	return h.db.Close()
}
