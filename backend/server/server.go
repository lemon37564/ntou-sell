package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"runtime"
	"se/backend/bid"
	"se/backend/cart"
	"se/backend/history"
	"se/backend/order"
	"se/backend/product"
	"se/backend/sell"
	"se/backend/user"

	"github.com/gorilla/mux"
)

// Server handle all services
type Server struct {
	DB *sql.DB
	Ur *user.User
	Pd *product.Product
	Od *order.Order
	Ht *history.History
	Bd *bid.Bid
	Ct *cart.Cart
	Se *sell.Sell
}

// Serve start all functions provided for user
func (ser *Server) Serve() {
	osys := runtime.GOOS
	log.Println("system:", osys)

	port := os.Getenv("PORT")

	// when test on localhost
	if osys == "windows" {
		port = "8080"
	}
	log.Println("Service running on port:", port)

	r := mux.NewRouter()

	r.HandleFunc("/backend/{key}", ser.defaultFunc)
	r.HandleFunc("/backend/bid/{key}", ser.fetchBid)
	r.HandleFunc("/backend/cart/{key}", ser.fetchCart)
	r.HandleFunc("/backend/history/{key}", ser.fetchHistory)
	r.HandleFunc("/backend/order/{key}", ser.fetchOrder)
	r.HandleFunc("/backend/product/{key}", ser.fetchProduct)
	r.HandleFunc("/backend/sell/{key}", ser.fetchSell)
	r.HandleFunc("/backend/user/{key}", ser.fetchUser)

	r.HandleFunc("/backend/pics/{key}", ser.picHandler)

	http.Handle("/backend/", r)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("webpage"))))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
