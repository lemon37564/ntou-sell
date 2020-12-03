package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"se/database"
	"se/user"
)

func main() {
	database.RemoveAll() // clear all the data in database

	db := database.Open()
	defer db.Close()

	database.TestInsert(db)
	database.TestSearch(db)

	newWeb := web{db: db}
	newWeb.weber()
}

//////////////////////////////////
//////////////////////////////////
///// snychorization problem /////
//////////////////////////////////
//////////////////////////////////

type web struct {
	db *sql.DB
}

func (we *web) weber() {
	http.HandleFunc("/", we.service)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (we *web) service(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Printf("connection<host: %v, remote: %v>\n", r.Host, r.RemoteAddr)

	u := user.NewUser(we.db)
	fmt.Fprintf(w, u.GetAllUserData())
}
