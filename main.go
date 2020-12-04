package main

import (
	"fmt"
	"log"
	"net/http"
	"se/database"
	"se/user"
)

func main() {
	database.RemoveAll() // clear all the data in database
	database.Check()

	database.TestInsert()
	// database.TestSearch()
	web()
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

	u := user.UserInit()
	fmt.Fprintf(w, u.GetAllUserData())
}
