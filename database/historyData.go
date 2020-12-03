package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const historyTable = `CREATE TABLE history(
						uid int NOT NULL,
						pd_id int NOT NULL,
						PRIMARY KEY(uid, pd_id),
						FOREIGN KEY(uid) REFERENCES user
					);`

type HistoryData struct {
	db *sql.DB

	insert  *sql.Stmt
	_delete *sql.Stmt
}

func HistoryDataInit() *HistoryData {
	history := new(HistoryData)

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	history.db = db

	history.insert, err = db.Prepare("INSERT INTO history VALUES(?,?);")
	if err != nil {
		log.Fatal(err)
	}

	history._delete, err = db.Prepare("DELETE FROM history where uid=? AND pd_id=?;")
	if err != nil {
		log.Fatal(err)
	}

	return history
}

// AddHistory add a single record into database
// return's an error
func (h *HistoryData) AddHistory(uid, pdid int) error {
	_, err := h.insert.Exec(uid, pdid)
	return err
}

func (h *HistoryData) Delete(uid, pdid int) error {
	_, err := h._delete.Exec(uid, pdid)
	return err
}

// WARNING: SQL injection
func (h *HistoryData) GetHistory(uid int) (pdid []int) {
	rows, err := h.db.Query("SELECT pd_id FROM history WHERE uid=" + fmt.Sprintf("%d", uid) + ";")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var d int
		err = rows.Scan(&d)
		if err != nil {
			log.Fatal(err)
		}

		pdid = append(pdid, d)
	}

	return
}

// always use this function at the end
func (h *HistoryData) DBClose() error {
	return h.db.Close()
}
