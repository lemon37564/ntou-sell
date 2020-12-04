package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// seq is for ordering
const historyTable = `CREATE TABLE history(
						uid int NOT NULL,
						pd_id int NOT NULL,
						seq int NOT NULL,
						PRIMARY KEY(uid, pd_id),
						FOREIGN KEY(uid) REFERENCES user
					);`

// HistoryDB contain funcions to use
type HistoryDB struct {
	insert  *sql.Stmt
	_delete *sql.Stmt
	delAll  *sql.Stmt
	maxSeq  *sql.Stmt
	get     *sql.Stmt
}

// HistoryDBInit prepare function for database using
func HistoryDBInit(db *sql.DB) *HistoryDB {
	var err error
	history := new(HistoryDB)

	history.insert, err = db.Prepare("INSERT INTO history VALUES(?,?,?);")
	if err != nil {
		panic(err)
	}

	history._delete, err = db.Prepare("DELETE FROM history where uid=? AND pd_id=?;")
	if err != nil {
		panic(err)
	}

	history.delAll, err = db.Prepare("DELETE FROM history WHERE uid=?;")
	if err != nil {
		panic(err)
	}

	history.maxSeq, err = db.Prepare("SELECT max(seq) FROM history WHERE uid=?;")
	if err != nil {
		panic(err)
	}

	history.get, err = db.Prepare("SELECT pd_id FROM history WHERE uid=? ORDER BY DESC seq;")
	if err != nil {
		panic(err)
	}

	return history
}

// AddHistory add a single record into database
// return's an error
// may encounter error when there's no history (beacuse max(seq) = null)
func (h *HistoryDB) AddHistory(uid, pdid int) error {
	rows, err := h.maxSeq.Query(uid)
	if err != nil {
		panic(err)
	}

	var seq int
	for rows.Next() {
		err = rows.Scan(&seq)
		if err != nil {
			panic(err)
		}
	}

	_, err = h.insert.Exec(uid, pdid, seq+1)
	return err
}

// Delete history with user id and product id
func (h *HistoryDB) Delete(uid, pdid int) error {
	_, err := h._delete.Exec(uid, pdid)
	return err
}

// DeleteAll deletes all history of a user by user id
func (h *HistoryDB) DeleteAll(uid int) error {
	_, err := h.delAll.Query()
	return err
}

// GetAll return all history of a user by id (descend order by time)
func (h *HistoryDB) GetAll(uid int) (pdid []int) {
	rows, err := h.get.Query(uid)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var d int
		err = rows.Scan(&d)
		if err != nil {
			panic(err)
		}

		pdid = append(pdid, d)
	}

	return
}
