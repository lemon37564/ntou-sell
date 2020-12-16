package main

import (
	"se/database"
	"se/server"
	"se/server/backend"
)

func main() {
	db := database.Open()
	defer db.Close()

	database.TestInsert(db)

	service := server.Server{
		DB: db,
		Ur: backend.UserInit(db),
		Pd: backend.ProductInit(db),
		Od: backend.OrderInit(db),
		Ht: backend.HistoryInit(db),
		Bd: backend.BidInit(db),
		Ct: backend.CartInit(db),
		Ms: backend.MessageInit(db)}
	service.Serve()
}
