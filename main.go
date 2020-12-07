package main

import (
	"se/bid"
	"se/database"
	"se/history"
	"se/order"
	"se/product"
	"se/user"
)

func main() {
	// database.RemoveAll() // clear all the data in database

	db := database.Open()
	defer db.Close()

	//database.TestInsert(db)
	// database.TestSearch(db)

	newWeb := server{
		db: db,
		u:  user.NewUser(db),
		p:  product.ProductInit(db),
		o:  order.NewOrder(db),
		h:  history.NewHistory(db),
		b:  bid.NewBid(db)}
	newWeb.weber()
}
