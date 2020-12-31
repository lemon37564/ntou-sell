package database

import (
	"database/sql"
	"log"
	"time"
)

// rename order as orders (order is a keword in SQL)
const ordersTable = `
CREATE TABLE IF NOT EXISTS orders(
	uid int NOT NULL,
	pd_id int NOT NULL,
	amount int,
	state varchar(8),
	order_date timestamp,
	PRIMARY KEY(uid, pd_id),
	FOREIGN KEY(uid) REFERENCES user ON DELETE CASCADE,
	FOREIGN KEY(pd_id) REFERENCES product ON DELETE CASCADE,
	CHECK (amount > 0)
);`

// Order type store data of a single order
type Order struct {
	Pdid   int
	PdName string
	Amount int
	Price  int
	State  string
	Time   time.Time
}

type orderStmt struct {
	add      *sql.Stmt
	del      *sql.Stmt
	upAmt    *sql.Stmt
	getAll   *sql.Stmt
	getOrder *sql.Stmt
	getPd    *sql.Stmt
}

func orderPrepare(db *sql.DB) *orderStmt {
	var err error
	order := new(orderStmt)

	const (
		add      = "INSERT INTO orders VALUES(?,?,?,?,?);"
		del      = "DELETE FROM orders WHERE uid=? AND pd_id=?;"
		upAmt    = "UPDATE orders SET amount=? WHERE uid=? AND pd_id=?;"
		getAll   = "SELECT pd_id, amount, state, order_date FROM orders WHERE uid=? ORDER BY order_date DESC;"
		getOrder = "SELECT pd_id, amount, state, order_date FROM orders WHERE uid=? AND pd_id=?;"
		getPd    = "SELECT product_name, price FROM product WHERE pd_id=? AND pd_id>0;"
	)

	if order.add, err = db.Prepare(add); err != nil {
		log.Println(err)
	}

	if order.del, err = db.Prepare(del); err != nil {
		log.Println(err)
	}

	if order.upAmt, err = db.Prepare(upAmt); err != nil {
		log.Println(err)
	}

	if order.getAll, err = db.Prepare(getAll); err != nil {
		log.Println(err)
	}

	if order.getOrder, err = db.Prepare(getOrder); err != nil {
		log.Println(err)
	}

	if order.getPd, err = db.Prepare(getPd); err != nil {
		log.Println(err)
	}

	return order
}

// AddOrder add order into order of specific user by user id
func (dt Data) AddOrder(uid, pdid, amount int, date time.Time) error {

	_, err := dt.Order.add.Exec(uid, pdid, amount, "unknown", date)
	return err
}

// DeleteOrder order by user id and product id
func (dt Data) DeleteOrder(uid, pdid int) error {
	_, err := dt.Order.del.Exec(uid, pdid)
	return err
}

// GetOrderByUIDAndPdid return order by user id and product id
func (dt Data) GetOrderByUIDAndPdid(uid, pdid int) Order {
	rows, err := dt.Order.getOrder.Query(uid, pdid)
	if err != nil {
		log.Println(err)
		return Order{}
	}

	var od Order
	for rows.Next() {
		err = rows.Scan(&od.Pdid, &od.Amount, &od.State, &od.Time)
		if err != nil {
			log.Println(err)
			return Order{}
		}
	}

	od.PdName, od.Price = dt.getPdNameAndPrice(pdid)

	return od
}

// GetAllOrder return all order with type Order, need argument user id
func (dt Data) GetAllOrder(uid int) (ods []Order) {
	rows, err := dt.Order.getAll.Query(uid)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var tmp Order
		err = rows.Scan(&tmp.Pdid, &tmp.Amount, &tmp.State, &tmp.Time)
		if err != nil {
			log.Println(err)
		}

		ods = append(ods, tmp)
	}

	for i, v := range ods {
		ods[i].PdName, ods[i].Price = dt.getPdNameAndPrice(v.Pdid)
	}

	return
}

func (dt Data) getPdNameAndPrice(pdid int) (string, int) {
	var name string
	var price int

	rows, err := dt.Order.getPd.Query(pdid)
	if err != nil {
		log.Println(err)
		return "", -1
	}

	for rows.Next() {
		err = rows.Scan(&name, &price)
		if err != nil {
			log.Println(err)
			return "", -1
		}
	}

	return name, price
}
