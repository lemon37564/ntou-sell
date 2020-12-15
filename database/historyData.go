package database

import (
	"database/sql"
	"log"

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
	insert        *sql.Stmt
	_delete       *sql.Stmt
	delAll        *sql.Stmt
	maxSeq        *sql.Stmt
	getnew        *sql.Stmt
	getold        *sql.Stmt
	getPdFrompdid *sql.Stmt
}

// History contain product ids
type History struct {
	Pdid int
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

	history.maxSeq, err = db.Prepare("SELECT max(seq) FROM history;")
	if err != nil {
		panic(err)
	}

	history.getnew, err = db.Prepare("SELECT pd_id FROM history WHERE uid=? ORDER BY seq DESC LIMIT ?;")
	if err != nil {
		panic(err)
	}

	history.getold, err = db.Prepare("SELECT pd_id FROM history WHERE uid=? ORDER BY seq ASC LIMIT ?;")
	if err != nil {
		panic(err)
	}

	history.getPdFrompdid, err = db.Prepare("SELECT * FROM product WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	return history
}

// AddHistory add a single record into database
// return's an error
// may encounter error when there's no history (beacuse max(seq) = null)
func (h *HistoryDB) AddHistory(uid, pdid int) error {
	// do this is to prevent history duplicate (delete the old one and add a new one)
	// then the new one will be close to the front.
	_, err := h._delete.Exec(uid, pdid)
	if err != nil {
		log.Println(err)
		return err
	}

	rows, err := h.maxSeq.Query()
	if err != nil {
		log.Println(err)
		return err
	}

	var seq int
	for rows.Next() {
		err = rows.Scan(&seq)
		if err != nil {
			log.Println(err)
			// error when user has no history (set seq as 0)
			seq = 0
		}
	}

	seq++

	_, err = h.insert.Exec(uid, pdid, seq)
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

// Get return all history of a user by id (descend order by time)
func (h *HistoryDB) Get(uid int, amount int, newest bool) (all []Product) {

	var rows *sql.Rows
	var err error
	var pdids []int

	if newest {
		rows, err = h.getnew.Query(uid, amount)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		rows, err = h.getold.Query(uid, amount)
		if err != nil {
			log.Println(err)
			return
		}
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return
		}

		pdids = append(pdids, id)
	}

	for _, v := range pdids {
		all = append(all, h.getPdByPdid(v))
	}

	return
}

func (h *HistoryDB) getPdByPdid(pdid int) (pd Product) {

	rows, err := h.getPdFrompdid.Query(pdid)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&pd.Pdid, &pd.PdName, &pd.Price, &pd.Description, &pd.Amount, &pd.Eval, &pd.SellerID, &pd.Bid, &pd.Date)
		if err != nil {
			log.Println(err)
			return
		}
	}

	return
}
