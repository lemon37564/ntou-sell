package database

import (
	"database/sql"
	"log"
	"time"
)

// rename order as orders (order is a keword in SQL)
const ordersTable = `CREATE TABLE IF NOT EXISTS orders(
						uid int NOT NULL,
						pd_id int NOT NULL,
						amount int,
						state varchar(8),
						order_date timestamp,
						PRIMARY KEY(uid, pd_id),
						FOREIGN KEY(uid) REFERENCES user,
						FOREIGN KEY(pd_id) REFERENCES product
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

// OrderDB contain funcions to use
type OrderDB struct {
	insert            *sql.Stmt
	_delete           *sql.Stmt
	updateAmount      *sql.Stmt
	getall            *sql.Stmt
	getOrder          *sql.Stmt
	getPdNameAndPrice *sql.Stmt
}

// OrderDBInit prepare function for database using
func OrderDBInit(db *sql.DB) *OrderDB {
	var err error
	order := new(OrderDB)

	order.insert, err = db.Prepare("INSERT INTO orders VALUES(?,?,?,?,?);")
	if err != nil {
		panic(err)
	}

	order._delete, err = db.Prepare("DELETE FROM orders WHERE uid=? AND pd_id=?;")
	if err != nil {
		panic(err)
	}

	order.updateAmount, err = db.Prepare("UPDATE orders SET amount=? WHERE uid=? AND pd_id=?;")
	if err != nil {
		panic(err)
	}

	order.getall, err = db.Prepare("SELECT pd_id, amount, state, order_date FROM orders WHERE uid=? ORDER BY order_date DESC;")
	if err != nil {
		panic(err)
	}

	order.getOrder, err = db.Prepare("SELECT pd_id, amount, state, order_date FROM orders WHERE uid=? AND pd_id=?;")
	if err != nil {
		panic(err)
	}

	order.getPdNameAndPrice, err = db.Prepare("SELECT product_name, price FROM product WHERE pd_id=? AND pd_id>0;")
	if err != nil {
		panic(err)
	}

	return order
}

// AddOrder add order into order of specific user by user id
func (o *OrderDB) AddOrder(uid, pdid, amount int, date time.Time) error {

	_, err := o.insert.Exec(uid, pdid, amount, "unknown", date)
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

	od.PdName, od.Price = o.getPdNameAndPriceByPdID(pdid)

	return od
}

// GetAllOrder return all order with type Order, need argument user id
func (o *OrderDB) GetAllOrder(uid int) (ods []Order) {
	rows, err := o.getall.Query(uid)
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
		ods[i].PdName, ods[i].Price = o.getPdNameAndPriceByPdID(v.Pdid)
	}

	return
}

func (o *OrderDB) getPdNameAndPriceByPdID(pdid int) (string, int) {
	var name string
	var price int

	rows, err := o.getPdNameAndPrice.Query(pdid)
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
