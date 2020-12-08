package database

import (
	"database/sql"
)

const bidTable = `CREATE TABLE bid(
					pd_id int NOT NULL,
					deadline varchar(16) NOT NULL,
					now_bidder_uid int NOT NULL,
					now_money int NOT NULL,
					seller_uid int NOT NULL,
					PRIMARY KEY(pd_id),
					FOREIGN KEY(seller_uid) REFERENCES user
					FOREIGN KEY(now_bidder_uid) REFERENCES user
				);`

// Bid struct store data of a single bid
type Bid struct {
	Deadline    string
	NowBidderID int
	NowMoney    int
	UID         int
}

// BidDB contain functions to use
type BidDB struct {
	insert         *sql.Stmt
	_delete        *sql.Stmt
	updateUID      *sql.Stmt
	updateMoney    *sql.Stmt
	updateDeadLine *sql.Stmt
	getAllBid      *sql.Stmt
	getBid         *sql.Stmt
}

// BidDataInit prepare functions for database using. require arg *sql.DB
func BidDataInit(db *sql.DB) *BidDB {
	var err error
	bid := new(BidDB)

	bid.insert, err = db.Prepare("INSERT INTO bid VALUES(?,?,?,?,?);")
	if err != nil {
		panic(err)
	}

	bid._delete, err = db.Prepare("DELETE FROM bid WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	bid.updateUID, err = db.Prepare("UPDATE bid SET seller_uid=? WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	bid.updateMoney, err = db.Prepare("UPDATE bid SET now_money=? WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	bid.updateDeadLine, err = db.Prepare("UPDATE bid SET deadline=? WHERE pd_id=?;")
	if err != nil {
		panic(err)
	}

	bid.getAllBid, err = db.Prepare("SELECT deadline, now_bidder_uid, now_money, seller_uid FROM bid")
	if err != nil {
		panic(err)
	}

	bid.getBid, err = db.Prepare("SELECT deadline, now_bidder_uid, now_money, seller_uid FROM bid WHERE pd_id=?")
	if err != nil {
		panic(err)
	}

	return bid
}

// AddNewBid insert new bid information into database
func (b *BidDB) AddNewBid(pdid int, deadline string, lowestMoney int, uid int) error {
	_, err := b.insert.Exec(pdid, deadline, nil, lowestMoney, uid)
	return err
}

// DeleteBid delete specific bid with pd_id
func (b *BidDB) DeleteBid(pdid int) error {
	_, err := b._delete.Exec(pdid)
	return err
}

// NewBidderGet update bidder_id and money if anyone won the bid at a time
func (b *BidDB) NewBidderGet(pdid, bidderID, money int) error {
	_, err := b.updateUID.Exec(bidderID, pdid)
	if err != nil {
		return err
	}

	_, err = b.updateMoney.Exec(money, pdid)
	return err
}

// GetBidByID return informations by product id
func (b *BidDB) GetBidByID(pdid int) (bid Bid) {
	rows, err := b.getBid.Query(pdid)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		bid = *new(Bid)
		err = rows.Scan(&bid.Deadline, &bid.NowBidderID, &bid.NowMoney, &bid.UID)
		if err != nil {
			panic(err)
		}
	}

	return
}

// GetAllBid return all bid product informations
func (b *BidDB) GetAllBid() (all []Bid) {
	rows, err := b.getAllBid.Query()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		bid := *new(Bid)
		err = rows.Scan(&bid.Deadline, &bid.NowBidderID, &bid.NowMoney, &bid.UID)
		if err != nil {
			panic(err)
		}

		all = append(all, bid)
	}

	return
}
