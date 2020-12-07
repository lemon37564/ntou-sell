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

	newserver := server{
		db: db,
		us: user.NewUser(db),
		pr: product.ProductInit(db),
		or: order.NewOrder(db),
		hi: history.NewHistory(db),
		bi: bid.NewBid(db)}
	newserver.server()
}
