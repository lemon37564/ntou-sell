package database

import (
	"database/sql"
	"log"
)

// seq is for ordering
const historyTable = `
CREATE TABLE IF NOT EXISTS history(
	uid int NOT NULL,
	pd_id int NOT NULL,
	seq int NOT NULL,
	PRIMARY KEY(uid, pd_id),
	FOREIGN KEY(uid) REFERENCES user
);`

// History contain product ids
type History struct {
	Pdid int
}

type historyStmt struct {
	add    *sql.Stmt
	del    *sql.Stmt
	delAll *sql.Stmt
	maxSeq *sql.Stmt
	getNew *sql.Stmt
	getOld *sql.Stmt
	getPd  *sql.Stmt
}

func historyPrepare(db *sql.DB) *historyStmt {
	var err error
	history := new(historyStmt)

	const (
		add    = "INSERT INTO history VALUES(?,?,?);"
		del    = "DELETE FROM history where uid=? AND pd_id=?;"
		delAll = "DELETE FROM history WHERE uid=?;"
		maxSeq = "SELECT max(seq) FROM history;"
		getNew = "SELECT pd_id FROM history WHERE uid=? ORDER BY seq DESC LIMIT ?;"
		getOld = "SELECT pd_id FROM history WHERE uid=? ORDER BY seq ASC LIMIT ?;"
		getPd  = "SELECT * FROM product WHERE pd_id=?;"
	)

	if history.add, err = db.Prepare(add); err != nil {
		log.Println(err)
	}

	if history.del, err = db.Prepare(del); err != nil {
		log.Println(err)
	}

	if history.delAll, err = db.Prepare(delAll); err != nil {
		log.Println(err)
	}

	if history.maxSeq, err = db.Prepare(maxSeq); err != nil {
		log.Println(err)
	}

	if history.getNew, err = db.Prepare(getNew); err != nil {
		log.Println(err)
	}

	if history.getOld, err = db.Prepare(getOld); err != nil {
		log.Println(err)
	}

	if history.getPd, err = db.Prepare(getPd); err != nil {
		log.Println(err)
	}

	return history
}

// AddHistory add a single record into database
// return's an error
// may encounter error when there's no history (beacuse max(seq) = null)
func (dt Data) AddHistory(uid, pdid int) error {
	// do this is to prevent history duplicate (delete the old one and add a new one)
	// then the new one will be close to the front.
	_, err := dt.history.del.Exec(uid, pdid)
	if err != nil {
		log.Println(err)
		return err
	}

	rows, err := dt.history.maxSeq.Query()
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

	_, err = dt.history.add.Exec(uid, pdid, seq)
	return err
}

// DeleteHistory with user id and product id
func (dt Data) DeleteHistory(uid, pdid int) error {
	_, err := dt.history.del.Exec(uid, pdid)
	return err
}

// DeleteAllHistory deletes all history of a user by user id
func (dt Data) DeleteAllHistory(uid int) error {
	_, err := dt.history.delAll.Query()
	return err
}

// GetAllHistory return all history of a user by id (descend order by time)
func (dt Data) GetAllHistory(uid int, amount int, newest bool) (all []Product) {
	var (
		rows  *sql.Rows
		err   error
		pdids []int
		id    int
	)

	if newest {
		rows, err = dt.history.getNew.Query(uid, amount)
	} else {
		rows, err = dt.history.getOld.Query(uid, amount)
	}

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return
		}

		pdids = append(pdids, id)
	}

	for _, v := range pdids {
		all = append(all, dt.getPdByPdid(v))
	}

	return
}

func (dt Data) getPdByPdid(pdid int) (pd Product) {

	rows, err := dt.history.getPd.Query(pdid)
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
