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
	FOREIGN KEY(seller_uid) REFERENCES user,
	FOREIGN KEY(now_bidder_uid) REFERENCES user
);`

// Bid struct store data of a single bid
type Bid struct {
	Deadline    time.Time
	NowBidderID int
	NowMoney    int
	UID         int
}

type bidStmt struct {
	add    *sql.Stmt
	del    *sql.Stmt
	upUID  *sql.Stmt
	upMon  *sql.Stmt
	upDL   *sql.Stmt
	getAll *sql.Stmt
	getBid *sql.Stmt
	getPd  *sql.Stmt
}

func bidPrepare(db *sql.DB) *bidStmt {
	var err error
	bid := new(bidStmt)

	const (
		add    = "INSERT INTO bid VALUES(?,?,?,?,?);"
		del    = "DELETE FROM bid WHERE pd_id=?;"
		upUID  = "UPDATE bid SET seller_uid=? WHERE pd_id=?;"
		upMon  = "UPDATE bid SET now_money=? WHERE pd_id=?;"
		upDL   = "UPDATE bid SET deadline=? WHERE pd_id=?;"
		getAll = "SELECT deadline, now_bidder_uid, now_money, seller_uid FROM bid;"
		getBid = "SELECT deadline, now_bidder_uid, now_money, seller_uid FROM bid WHERE pd_id=?;"
		getPd  = "SELECT * FROM bid WHERE pd_id=?;"
	)

	if bid.add, err = db.Prepare(add); err != nil {
		log.Println(err)
	}

	if bid.del, err = db.Prepare(del); err != nil {
		log.Println(err)
	}

	if bid.upUID, err = db.Prepare(upUID); err != nil {
		log.Println(err)
	}

	if bid.upMon, err = db.Prepare(upMon); err != nil {
		log.Println(err)
	}

	if bid.upDL, err = db.Prepare(upDL); err != nil {
		log.Println(err)
	}

	if bid.getAll, err = db.Prepare(getAll); err != nil {
		log.Println(err)
	}

	if bid.getBid, err = db.Prepare(getBid); err != nil {
		log.Println(err)
	}

	if bid.getPd, err = db.Prepare(getPd); err != nil {
		log.Println(err)
	}

	return bid
}

// AddBid add new bid information into database
func (dt Data) AddBid(pdid int, deadline time.Time, lowestMoney int, uid int) error {
	_, err := dt.bid.add.Exec(pdid, deadline, nil, lowestMoney, uid)
	return err
}

// DeleteBid delete specific bid with pd_id
func (dt Data) DeleteBid(pdid int) error {
	_, err := dt.bid.del.Exec(pdid)
	return err
}

// WonBid update bidder_id and money if anyone won the bid at a time
func (dt Data) WonBid(pdid, bidderID, money int) error {
	_, err := dt.bid.upUID.Exec(bidderID, pdid)
	if err != nil {
		return err
	}

	_, err = dt.bid.upMon.Exec(money, pdid)
	return err
}

// GetBidByID return informations by product id
func (dt Data) GetBidByID(pdid int) (bid Bid) {
	rows, err := dt.bid.getBid.Query(pdid)
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
func (dt Data) GetAllBidProducts(pdid int) (pd Product) {
	rows, err := dt.bid.getBid.Query(pdid)
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
func (dt Data) GetAllBid() (all []Bid) {
	rows, err := dt.bid.getAll.Query()
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
