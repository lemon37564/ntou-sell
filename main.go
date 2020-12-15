package main

import (
	"se/database"
	"se/server"
	"se/server/backend"
	"time"
)

func main() {
	db := database.Open()
	defer db.Close()

	database.TestInsert(db)

	list := make(map[string]int)
	blist := make(map[string]bool)

	service := server.Server{
		DB:        db,
		Ur:        backend.NewUser(db),
		Pd:        backend.ProductInit(db),
		Od:        backend.NewOrder(db),
		Ht:        backend.NewHistory(db),
		Bd:        backend.NewBid(db),
		Ct:        backend.NewCart(db),
		Se:        backend.NewSell(db),
		Ms:        backend.NewMessage(db),
		Timer:     time.Now(),
		IPList:    list,
		BlackList: blist}
	service.Serve()
}
