package database

import (
	"database/sql"
	"log"
	"time"
)

const bidTable = `
CREATE TABLE IF NOT EXISTS bid(
	pd_id int NOT NULL,
	deadline timestamp NOT NULL,
	now_bidder_uid int NOT NULL,
	now_money int NOT NULL,
	seller_uid int NOT NULL,
	PRIMARY KEY(pd_id),
	FOREIGN KEY(seller_uid) REFERENCES user ON DELETE CASCADE,
	FOREIGN KEY(now_bidder_uid) REFERENCES user ON DELETE CASCADE,
	CHECK (now_money > 0)
);`

// Bid struct store data of a single bid
type Bid struct {
	Deadline    time.Time
	NowBidderID int
	NowMoney    int
	UID         int
}

var (
	bidAdd    *sql.Stmt
	bidDel    *sql.Stmt
	bidWon    *sql.Stmt
	bidUpDL   *sql.Stmt
	bidGetAll *sql.Stmt
	bidGet    *sql.Stmt
	bidGetPd  *sql.Stmt
)

func bidPrepare(db *sql.DB) {
	var err error

	const (
		add    = "INSERT INTO bid VALUES(?,?,?,?,?);"
		del    = "DELETE FROM bid WHERE pd_id=?;"
		won    = "UPDATE product SET price=? WHERE pd_id=?;"
		upDL   = "UPDATE bid SET deadline=? WHERE pd_id=?;"
		getAll = "SELECT deadline, now_bidder_uid, now_money, seller_uid FROM bid;"
		getBid = "SELECT deadline, now_bidder_uid, now_money, seller_uid FROM bid WHERE pd_id=?;"
		getPd  = "SELECT * FROM bid WHERE pd_id=?;"
	)

	if bidAdd, err = db.Prepare(add); err != nil {
		log.Println(err)
	}

	if bidDel, err = db.Prepare(del); err != nil {
		log.Println(err)
	}

	if bidWon, err = db.Prepare(won); err != nil {
		log.Println(err)
	}

	if bidUpDL, err = db.Prepare(upDL); err != nil {
		log.Println(err)
	}

	if bidGetAll, err = db.Prepare(getAll); err != nil {
		log.Println(err)
	}

	if bidGet, err = db.Prepare(getBid); err != nil {
		log.Println(err)
	}

	if bidGetPd, err = db.Prepare(getPd); err != nil {
		log.Println(err)
	}
}

// AddBid add new bid information into database
func AddBid(pdid int, deadline time.Time, lowestMoney int, uid int) error {
	_, err := bidAdd.Exec(pdid, deadline, nil, lowestMoney, uid)
	return err
}

// DeleteBid delete specific bid with pd_id
func DeleteBid(pdid int) error {
	_, err := bidDel.Exec(pdid)
	return err
}

// WonBid update bidder_id and money if anyone won the bid at a time
func WonBid(uid, pdid, money int) error {
	_, err := bidWon.Exec(money, pdid)
	return err
}

// GetBidByID return informations by product id
func GetBidByID(pdid int) (bid Bid) {
	rows, err := bidGet.Query(pdid)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&bid.Deadline, &bid.NowBidderID, &bid.NowMoney, &bid.UID)
		if err != nil {
			log.Println(err)
		}
	}

	return
}

// GetAllBidProducts return literally product info by pdid
func GetAllBidProducts(pdid int) (pd Product) {
	rows, err := bidGet.Query(pdid)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&pd.Pdid, &pd.PdName, &pd.Price, &pd.Description, &pd.Amount, &pd.Eval, &pd.SellerID, &pd.Bid, &pd.Date)
		if err != nil {
			log.Println(err)
		}
	}

	return
}

// GetAllBid return all bid product informations
func GetAllBid() (all []Bid) {
	rows, err := bidGetAll.Query()
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		bid := *new(Bid)
		err = rows.Scan(&bid.Deadline, &bid.NowBidderID, &bid.NowMoney, &bid.UID)
		if err != nil {
			log.Println(err)
		}

		all = append(all, bid)
	}

	return
}
