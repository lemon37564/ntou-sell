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
	FOREIGN KEY(uid) REFERENCES userDB ON DELETE CASCADE,
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

var (
	ordAdd      *sql.Stmt
	ordDel      *sql.Stmt
	ordUpAmt    *sql.Stmt
	ordGetAll   *sql.Stmt
	ordGetOrder *sql.Stmt
	ordGetPd    *sql.Stmt
)

func orderPrepare(db *sql.DB) {
	var err error

	const (
		add      = "INSERT INTO orders VALUES(?,?,?,?,?);"
		del      = "DELETE FROM orders WHERE uid=? AND pd_id=?;"
		upAmt    = "UPDATE orders SET amount=? WHERE uid=? AND pd_id=?;"
		getAll   = "SELECT pd_id, amount, state, order_date FROM orders WHERE uid=? ORDER BY order_date DESC;"
		getOrder = "SELECT pd_id, amount, state, order_date FROM orders WHERE uid=? AND pd_id=?;"
		getPd    = "SELECT product_name, price FROM product WHERE pd_id=? AND pd_id>0;"
	)

	if ordAdd, err = db.Prepare(add); err != nil {
		panic(err)
	}

	if ordDel, err = db.Prepare(del); err != nil {
		panic(err)
	}

	if ordUpAmt, err = db.Prepare(upAmt); err != nil {
		panic(err)
	}

	if ordGetAll, err = db.Prepare(getAll); err != nil {
		panic(err)
	}

	if ordGetOrder, err = db.Prepare(getOrder); err != nil {
		panic(err)
	}

	if ordGetPd, err = db.Prepare(getPd); err != nil {
		panic(err)
	}
}

// AddOrder add order into order of specific user by user id
func AddOrder(uid, pdid, amount int, date time.Time) error {

	_, err := ordAdd.Exec(uid, pdid, amount, "unknown", date)
	return err
}

// DeleteOrder order by user id and product id
func DeleteOrder(uid, pdid int) error {
	_, err := ordDel.Exec(uid, pdid)
	return err
}

// GetOrderByUIDAndPdid return order by user id and product id
func GetOrderByUIDAndPdid(uid, pdid int) Order {
	rows, err := ordGetOrder.Query(uid, pdid)
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

	rows.Close()
	od.PdName, od.Price = getPdNameAndPrice(pdid)

	return od
}

// GetAllOrder return all order with type Order, need argument user id
func GetAllOrder(uid int) (ods []Order) {
	rows, err := ordGetAll.Query(uid)
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
	rows.Close()

	for i, v := range ods {
		ods[i].PdName, ods[i].Price = getPdNameAndPrice(v.Pdid)
	}

	return
}

func getPdNameAndPrice(pdid int) (string, int) {
	var name string
	var price int

	rows, err := ordGetPd.Query(pdid)
	if err != nil {
		log.Println(err)
		return "", -1
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&name, &price)
		if err != nil {
			log.Println(err)
			return "", -1
		}
	}

	return name, price
}
