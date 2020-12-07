package main

import (
	"database/sql"
	"log"
	"net/http"
	"se/bid"
	"se/history"
	"se/order"
	"se/product"
	"se/user"
)

type server struct {
	db *sql.DB
	us *user.User
	pr *product.Product
	or *order.Order
	hi *history.History
	bi *bid.Bid
}

func (ser *server) server() {
	http.HandleFunc("/", ser.service)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (ser *server) service(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	path := r.URL.Path
	query := r.URL.Query()

	log.Printf("< host: %v, remote: %v> path: %v, args: %v\n", r.Host, r.RemoteAddr, path, query)

	arg := path[1:] // eliminate first "/"

	ser.fetch(w, r, arg, query)
}
