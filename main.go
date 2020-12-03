package main

import (
	"se/database"
)

func main() {
	database.RemoveAll() // clear all the data in database
	database.Check()

	database.TestInsert()
	database.TestSearch()
}
