package main

import (
	"se/database"
	"se/server"
	"se/server/backend"
)

func main() {
	panic("stop")

	db := database.Open()
	defer db.Close()

	database.TestInsert(db)

	service := server.Server{
		DB: db,
		Ur: backend.NewUser(db),
		Pd: backend.ProductInit(db),
		Od: backend.NewOrder(db),
		Ht: backend.NewHistory(db),
		Bd: backend.NewBid(db),
		Ct: backend.NewCart(db),
		Se: backend.NewSell(db),
		Ms: backend.NewMessage(db)}
	service.Serve()
}
