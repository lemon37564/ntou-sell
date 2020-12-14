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
	insert  *sql.Stmt
	_delete *sql.Stmt
	delAll  *sql.Stmt
	maxSeq  *sql.Stmt
	get     *sql.Stmt
	getall  *sql.Stmt
}

type History struct {
	UID  int
	Pdid int
	Seq  int
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

	history.get, err = db.Prepare("SELECT * FROM product WHERE pd_id IN (SELECT pd_id FROM history WHERE uid=? ORDER BY seq DESC LIMIT ?);")
	if err != nil {
		panic(err)
	}

	history.getall, err = db.Prepare("SELECT * FROM history ORDER BY seq DESC;")
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
		log.Println(err)
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
func (h *HistoryDB) Get(uid int, amount int) (all []Product) {
	rows, err := h.get.Query(uid, amount)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var pd Product
		err = rows.Scan(&pd.Pdid, &pd.PdName, &pd.Price, &pd.Description, &pd.Amount, &pd.Eval, &pd.SellerID, &pd.Bid, &pd.Date)
		if err != nil {
			panic(err)
		}

		all = append(all, pd)
	}

	return
}

// GetAll return all historys(debugging only)
func (h *HistoryDB) GetAll() (all []History) {
	rows, err := h.getall.Query()
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var hi History
		err = rows.Scan(&hi.UID, &hi.Pdid, &hi.Seq)
		if err != nil {
			log.Println(err)
		}

		all = append(all, hi)
	}

	return
}
