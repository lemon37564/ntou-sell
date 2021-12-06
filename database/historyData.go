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
	FOREIGN KEY(uid) REFERENCES userDB ON DELETE CASCADE
);`

// History contain product ids
type History struct {
	Pdid int
}

var (
	histAdd    *sql.Stmt
	histDel    *sql.Stmt
	histDelAll *sql.Stmt
	histMaxSeq *sql.Stmt
	histGetNew *sql.Stmt
	histGetOld *sql.Stmt
	histGetPd  *sql.Stmt
)

func historyPrepare(db *sql.DB) {
	var err error

	const (
		add    = "INSERT INTO history VALUES($1,$2,$3);"
		del    = "DELETE FROM history where uid=$1 AND pd_id=$2;"
		delAll = "DELETE FROM history WHERE uid=$1;"
		maxSeq = "SELECT max(seq) FROM history;"
		getNew = "SELECT pd_id FROM history WHERE uid=$1 ORDER BY seq DESC LIMIT $2;"
		getOld = "SELECT pd_id FROM history WHERE uid=$1 ORDER BY seq ASC LIMIT $2;"
		getPd  = "SELECT * FROM product WHERE pd_id=$1;"
	)

	if histAdd, err = db.Prepare(add); err != nil {
		panic(err)
	}

	if histDel, err = db.Prepare(del); err != nil {
		panic(err)
	}

	if histDelAll, err = db.Prepare(delAll); err != nil {
		panic(err)
	}

	if histMaxSeq, err = db.Prepare(maxSeq); err != nil {
		panic(err)
	}

	if histGetNew, err = db.Prepare(getNew); err != nil {
		panic(err)
	}

	if histGetOld, err = db.Prepare(getOld); err != nil {
		panic(err)
	}

	if histGetPd, err = db.Prepare(getPd); err != nil {
		panic(err)
	}
}

// AddHistory add a single record into database
// return's an error
// may encounter error when there's no history (beacuse max(seq) = null)
func AddHistory(uid, pdid int) error {
	// do this is to prevent history duplicate (delete the old one and add a new one)
	// then the new one will be close to the front.
	_, err := histDel.Exec(uid, pdid)
	if err != nil {
		log.Println(err)
		return err
	}

	rows, err := histMaxSeq.Query()
	if err != nil {
		log.Println(err)
		return err
	}

	var seq int
	if rows.Next() {
		err = rows.Scan(&seq)
		if err != nil {
			log.Println(err)
			// error when user has no history (set seq as 0)
			seq = 0
		}
	}
	rows.Close()

	_, err = histAdd.Exec(uid, pdid, seq+1)
	return err
}

// DeleteHistory with user id and product id
func DeleteHistory(uid, pdid int) error {
	_, err := histDel.Exec(uid, pdid)
	return err
}

// DeleteAllHistory deletes all history of a user by user id
func DeleteAllHistory(uid int) error {
	_, err := histDelAll.Query()
	return err
}

// GetAllHistory return all history of a user by id (descend order by time)
func GetAllHistory(uid int, amount int, newest bool) (all []Product) {
	var (
		rows  *sql.Rows
		err   error
		pdids []int
		id    int
	)

	if newest {
		rows, err = histGetNew.Query(uid, amount)
	} else {
		rows, err = histGetOld.Query(uid, amount)
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
	rows.Close()

	for _, v := range pdids {
		all = append(all, getPdByPdid(v))
	}

	return
}

func getPdByPdid(pdid int) (pd Product) {

	rows, err := histGetPd.Query(pdid)
	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&pd.Pdid, &pd.PdName, &pd.Price, &pd.Description, &pd.Amount, &pd.Eval, &pd.SellerID, &pd.Bid, &pd.Date)
		if err != nil {
			log.Println(err)
			return
		}
	}

	return
}
