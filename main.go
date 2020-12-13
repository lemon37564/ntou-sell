package main

import (
	"se/backend/bid"
	"se/backend/cart"
	"se/backend/history"
	"se/backend/order"
	"se/backend/product"
	"se/backend/sell"
	"se/backend/server"
	"se/backend/user"
	"se/database"
)

func main() {
	db := database.Open()
	defer db.Close()

	database.TestInsert(db)

	service := server.Server{
		DB: db,
		Ur: user.NewUser(db),
		Pd: product.ProductInit(db),
		Od: order.NewOrder(db),
		Ht: history.NewHistory(db),
		Bd: bid.NewBid(db),
		Ct: cart.NewCart(db),
		Se: sell.NewSell(db)}
	service.Serve()
}
