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
	db := database.Open()
	defer db.Close()

	database.TestInsert(db)

	service := server{
		db: db,
		us: user.NewUser(db),
		pr: product.ProductInit(db),
		or: order.NewOrder(db),
		hi: history.NewHistory(db),
		bi: bid.NewBid(db)}
	service.server()
}
