package database

import (
	"database/sql"
	"log"
)

// rename order as orders (order is a keword in SQL)
const ordersTable = `CREATE TABLE orders(
						uid int NOT NULL,
						pd_id int NOT NULL,
						name varchar(256),
						price int,
						amount int,
						sum int,
						seller_uid int NOT NULL,
						state varchar(8),
						PRIMARY KEY(uid, pd_id),
						FOREIGN KEY(uid) REFERENCES user,
						FOREIGN KEY(seller_uid) REFERENCES user,
						FOREIGN KEY(pd_id) REFERENCES product
					);`

type OrderData struct {
	db *sql.DB

	insert  *sql.Stmt
	_delete *sql.Stmt
	update  *sql.Stmt
	_select *sql.Stmt
}

func OrderDataInit() *OrderData {
	order := new(OrderData)

	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	order.db = db

	insert, err := db.Prepare("INSERT INTO order values(?,?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	order.insert = insert

	_delete, err := db.Prepare("DELETE FROM order where pd_id=?;")
	if err != nil {
		log.Fatal(err)
	}
	order._delete = _delete

	update, err := db.Prepare("UPDATE order SET ?=?;")
	if err != nil {
		log.Fatal(err)
	}
	order.update = update

	_select, err := db.Prepare("SELECT * FROM order WHERE ?=?;")
	if err != nil {
		log.Fatal(err)
	}
	order._select = _select

	return order
}

// wait for implementation
func (o *OrderData) Insert() error {
	_, err := o.insert.Exec()
	return err
}

// wait for implementation
func (o *OrderData) Delete(pdid string) error {
	_, err := o._delete.Exec(pdid)
	return err
}

// wait for implementation
func (o *OrderData) Update(products string) error {
	return nil
}

// wait for implementation
func (o *OrderData) Select() (string, error) {
	return "", nil
}

func (o *OrderData) DBClose() error {
	return o.db.Close()
}
