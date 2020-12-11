package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"se/bid"
	"se/cart"
	"se/history"
	"se/order"
	"se/product"
	"se/sell"
	"se/user"
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

	Sess *session
}

// Serve start all functions provided for user
func (ser *Server) Serve() {
	port := os.Getenv("PORT")
	log.Println("Service running on port:", port)

	http.HandleFunc("/", ser.help)
	http.HandleFunc("/help", ser.help)
	http.HandleFunc("/bid", ser.fetchBid)
	http.HandleFunc("/cart", ser.fetchCart)
	http.HandleFunc("/history", ser.fetchHistory)
	http.HandleFunc("/order", ser.fetchOrder)
	http.HandleFunc("/product", ser.fetchProduct)
	http.HandleFunc("/sell", ser.fetchSell)
	http.HandleFunc("/user", ser.fetchUser)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
