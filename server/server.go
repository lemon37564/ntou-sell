package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"runtime"
	"se/bid"
	"se/cart"
	"se/history"
	"se/order"
	"se/product"
	"se/sell"
	"se/user"

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

	r.HandleFunc("/default/{key}", ser.defaultFunc)
	r.HandleFunc("/bid/{key}", ser.fetchBid)
	r.HandleFunc("/cart/{key}", ser.fetchCart)
	r.HandleFunc("/history/{key}", ser.fetchHistory)
	r.HandleFunc("/order/{key}", ser.fetchOrder)
	r.HandleFunc("/product/{key}", ser.fetchProduct)
	r.HandleFunc("/sell/{key}", ser.fetchSell)
	r.HandleFunc("/user/{key}", ser.fetchUser)

	r.HandleFunc("/pics/{key}", ser.picHandler)

	http.Handle("/", r)
	http.Handle("/ntou-sell", http.FileServer(http.Dir("/web")))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
