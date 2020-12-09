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
	"se/user"
)

type Server struct {
	DB *sql.DB
	Ur *user.User
	Pd *product.Product
	Od *order.Order
	Ht *history.History
	Bd *bid.Bid
	Ct *cart.Cart
}

// Serve start all functions provided for user
func (ser *Server) Serve() {
	port := os.Getenv("PORT")
	log.Println("Service running on port:", port)

	http.HandleFunc("/", ser.service)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (ser *Server) service(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	path := r.URL.Path
	query := r.URL.Query()

	log.Printf("<host: %v, remote: %v>path: %v, args: %v\n", r.Host, r.RemoteAddr, path, query)

	arg := path[1:] // eliminate first "/"

	ser.fetch(w, r, arg, query)
}

// verify if user is legel by using cookies
func (ser *Server) verify(w http.ResponseWriter, r *http.Request) bool {
	cookie, cookie2, exist := ser.getCookies(w, r)

	if !exist {
		return false
	}

	account := cookie.Value
	password := cookie2.Value

	return ser.Ur.Login(account, password)
}
