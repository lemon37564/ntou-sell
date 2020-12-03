package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const bidTable = `CREATE TABLE bid(
					pd_id int NOT NULL,
					deadline varchar(16) NOT NULL,
					now_bidder_uid int NOT NULL,
					now_money int NOT NULL,
					uid int NOT NULL,
					PRIMARY KEY(pd_id),
					FOREIGN KEY(uid) REFERENCES user
					FOREIGN KEY(now_bidder_uid) REFERENCES user
				);`

type BidData struct {
	db *sql.DB

	insert         *sql.Stmt
	_delete        *sql.Stmt
	updateUid      *sql.Stmt
	updateMoney    *sql.Stmt
	updateDeadLine *sql.Stmt
}

type Bid struct {
	Deadline    string
	NowBidderId int
	NowMoney    int
	Uid         int
}

func BidDataInit() *BidData {
	bid := new(BidData)

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	bid.db = db

	bid.insert, err = db.Prepare("INSERT INTO bid VALUES(?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	bid._delete, err = db.Prepare("DELETE FROM bid WHERE pd_id=?;")
	if err != nil {
		log.Fatal(err)
	}

	bid.updateUid, err = db.Prepare("UPDATE bid SET uid=? WHERE pd_id=?;")
	if err != nil {
		log.Fatal(err)
	}

	bid.updateMoney, err = db.Prepare("UPDATE bid SET now_money=? WHERE pd_id=?;")
	if err != nil {
		log.Fatal(err)
	}

	bid.updateDeadLine, err = db.Prepare("UPDATE bid SET deadline=? WHERE pd_id=?;")
	if err != nil {
		log.Fatal(err)
	}

	return bid
}

func (b *BidData) AddNewBid(pdid int, deadline string, lowest_money int, uid int) error {
	_, err := b.insert.Exec(pdid, deadline, nil, lowest_money, uid)
	return err
}

func (b *BidData) DeleteBid(pdid int) error {
	_, err := b._delete.Exec(pdid)
	return err
}

// NewBidder update bidder_id and money if anyone won the price
func (b *BidData) NewBidderGet(pdid, bidderId, money int) error {
	_, err := b.updateUid.Exec(bidderId, pdid)
	if err != nil {
		return err
	}

	_, err = b.updateMoney.Exec(money, pdid)
	return err
}

// wait for implementation
func (b *BidData) GetAllBid() (all []Bid) {
	rows, err := b.db.Query("SELECT deadline, now_bidder_id, now_money, uid FROM bid")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		bid := *new(Bid)
		err = rows.Scan(&bid.Deadline, &bid.NowBidderId, &bid.NowMoney, &bid.Uid)
		if err != nil {
			log.Fatal(err)
		}

		all = append(all, bid)
	}

	return
}

// always use this function at the end
func (b *BidData) DBClose() error {
	return b.db.Close()
}
