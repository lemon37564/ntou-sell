package main

import (
	"se/database"
	"se/user"
)

func main() {
	database.RemoveAll() // clear all the data in database

	db := database.Open()
	defer db.Close()

	database.TestInsert(db)
	//database.TestSearch(db)

	newWeb := server{db: db, u: user.NewUser(db)}
	newWeb.weber()
}
