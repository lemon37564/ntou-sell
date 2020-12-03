package main

import (
	"log"
	"net/http"
	"se/database"
)

func main() {
	database.RemoveAll() // clear all the data in database

	db := database.Open()
	defer db.Close()
	// u := database.UserDataInit(db)

	database.TestInsert(db)
	database.TestSearch(db)
}

func web() {
	http.HandleFunc("/", service)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func service(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Printf("connection<host: %v, remote: %v>\n", r.Host, r.RemoteAddr)

	// u := user.UserInit()
	// fmt.Fprintf(w, u.GetAllUserData())
}
