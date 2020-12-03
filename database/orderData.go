package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// rename order as orders (order is a keword in SQL)
const ordersTable = `CREATE TABLE orders(
						uid int NOT NULL,
						pd_id int NOT NULL,
						amount int,
						state varchar(8),
						PRIMARY KEY(uid, pd_id),
						FOREIGN KEY(uid) REFERENCES user,
						FOREIGN KEY(pd_id) REFERENCES product
					);`

type OrderData struct {
	db *sql.DB

	insert       *sql.Stmt
	_delete      *sql.Stmt
	updateAmount *sql.Stmt
}

func OrderDataInit() *OrderData {
	order := new(OrderData)

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	order.db = db

	order.insert, err = db.Prepare("INSERT INTO order values(?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	order._delete, err = db.Prepare("DELETE FROM order where uid=? and pd_id=?;")
	if err != nil {
		log.Fatal(err)
	}

	order.updateAmount, err = db.Prepare("UPDATE order SET amount=? WHERE uid=? AND pd_id=?")
	if err != nil {
		log.Fatal(err)
	}

	return order
}

func (o *OrderData) AddOrder(uid, pdid, amount int, state string) error {
	_, err := o.insert.Exec(uid, pdid, amount, state)
	return err
}

func (o *OrderData) Delete(uid, pdid int) error {
	_, err := o._delete.Exec(uid, pdid)
	return err
}

func (o *OrderData) UpdateAmount(uid, pdid, amount int) error {
	_, err := o.updateAmount.Exec(amount, uid, pdid)
	return err
}

func (o *OrderData) GetAllOrder(uid int) (pdid []int) {
	rows, err := o.db.Query("SELECT pd_id FROM order WHERE uid=" + fmt.Sprintf("%d", uid) + ";")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var p int
		err = rows.Scan(&p)
		if err != nil {
			log.Fatal(err)
		}

		pdid = append(pdid, p)
	}

	return
}

// always use this function at the end
func (o *OrderData) DBClose() error {
	return o.db.Close()
}
