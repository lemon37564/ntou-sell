package database

import (
	"database/sql"
)

// CREATE TABLE orders(
// 	id varchar(16) NOT NULL,
// 	pd_id varchar(16) NOT NULL,
// 	name varchar(256),
// 	price int,
// 	amount int,
// 	sum int,
// 	seller_id varchar(16) NOT NULL,
// 	state varchar(8),
// 	PRIMARY KEY(id, pd_id),
// 	FOREIGN KEY(id) REFERENCES user,
// 	FOREIGN KEY(seller_id) REFERENCES user,
// 	FOREIGN KEY(pd_id) REFERENCES product
// );

type OrderData struct {
	db *sql.DB

	insert  *sql.Stmt
	_delete *sql.Stmt
	update  *sql.Stmt
	_select *sql.Stmt
}

func OrderDataInit() (*OrderData, error) {
	order := new(OrderData)

	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return order, err
	}
	defer db.Close()
	order.db = db

	insert, err := db.Prepare("INSERT INTO order values(?,?,?,?,?,?,?,?,?);")
	if err != nil {
		return order, err
	}
	order.insert = insert

	_delete, err := db.Prepare("DELETE FROM order where pd_id=?;")
	if err != nil {
		return order, err
	}
	order._delete = _delete

	update, err := db.Prepare("UPDATE order SET ?=?;")
	if err != nil {
		return order, err
	}
	order.update = update

	_select, err := db.Prepare("SELECT * FROM order WHERE ?=?;")
	if err != nil {
		return order, err
	}
	order._select = _select

	return order, nil
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
