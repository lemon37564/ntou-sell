package main

import (
	"se/bid"
	"se/database"
	"se/history"
	"se/order"
	"se/product"
	"se/server"
	"se/user"
)

func main() {
	db := database.Open()
	defer db.Close()

	database.TestInsert(db)

	service := server.Server{
		DB:   db,
		Ur:   user.NewUser(db),
		Pd:   product.ProductInit(db),
		Od:   order.NewOrder(db),
		Ht:   history.NewHistory(db),
		Bd:   bid.NewBid(db),
		Sess: server.NewSession()}
	service.Serve()
}
