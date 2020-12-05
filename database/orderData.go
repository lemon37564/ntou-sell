package database

import (
	"database/sql"
	"fmt"

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

// Order type store data of a single order
type Order struct {
	Pdid   int
	Amount int
	State  string
}

func (o Order) String() (res string) {
	res += " product ID:       " + fmt.Sprintf("%d\n", o.Pdid) + "\n"
	res += " product Amount:   " + fmt.Sprintf("%d\n", o.Amount) + "\n"
	res += "product State:     " + fmt.Sprintf("%s\n", o.State) + "\n"
	return
}

// OrderDB contain funcions to use
type OrderDB struct {
	insert       *sql.Stmt
	_delete      *sql.Stmt
	updateAmount *sql.Stmt
	getall       *sql.Stmt
	getOrder     *sql.Stmt
}

// OrderDBInit prepare function for database using
func OrderDBInit(db *sql.DB) *OrderDB {
	var err error
	order := new(OrderDB)

	order.insert, err = db.Prepare("INSERT INTO order VALUES(?,?,?,?);")
	if err != nil {
		panic(err)
	}

	order._delete, err = db.Prepare("DELETE FROM order WHERE uid=? AND pd_id=?;")
	if err != nil {
		panic(err)
	}

	order.updateAmount, err = db.Prepare("UPDATE order SET amount=? WHERE uid=? AND pd_id=?;")
	if err != nil {
		panic(err)
	}

	order.getall, err = db.Prepare("SELECT pd_id, amount, state FROM order WHERE uid=?;")
	if err != nil {
		panic(err)
	}

	order.getOrder, err = db.Prepare("SELECT pd_id, amount, state FROM order WHERE uid=? AND pd_id=?;")
	if err != nil {
		panic(err)
	}

	return order
}

// AddOrder add order into order of specific user by user id
func (o *OrderDB) AddOrder(uid, pdid, amount int) error {
	_, err := o.insert.Exec(uid, pdid, amount, "unknown")
	return err
}

// Delete order by user id and product id
func (o *OrderDB) Delete(uid, pdid int) error {
	_, err := o._delete.Exec(uid, pdid)
	return err
}

// GetOrderByUIDAndPdid return order by user id and product id
func (o *OrderDB) GetOrderByUIDAndPdid(uid, pdid int) Order {
	rows, err := o.getOrder.Query(uid, pdid)
	if err != nil {
		panic(err)
	}

	var od Order
	for rows.Next() {
		err = rows.Scan(&od.Pdid, &od.Amount, &od.State)
		if err != nil {
			panic(err)
		}
	}

	return od
}

// GetAllOrder return all order with type Order, need argument user id
func (o *OrderDB) GetAllOrder(uid int) (ods []Order) {
	rows, err := o.getall.Query(uid)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var tmp Order
		err = rows.Scan(&tmp.Pdid, &tmp.Amount, &tmp.State)
		if err != nil {
			panic(err)
		}

		ods = append(ods, tmp)
	}

	return
}
