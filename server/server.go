package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"runtime"
	"se/server/backend"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

const (
	limitAccess = 60
	refreshTime = time.Second * 30
)

// Server handle all services
type Server struct {
	DB *sql.DB
	Ur *backend.User
	Pd *backend.Product
	Od *backend.Order
	Ht *backend.History
	Bd *backend.Bid
	Ct *backend.Cart
	Se *backend.Sell
	Ms *backend.Message

	Timer     time.Time
	IPList    map[string]int
	BlackList map[string]bool
	Lock      sync.Mutex
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
	r.HandleFunc("/backend/message/{key}", ser.fetchMessage)

	r.HandleFunc("/backend/pics/{key}", ser.picHandler)

	http.Handle("/backend/", r)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("webpage"))))

	go ser.refresh()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (ser *Server) validation(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()

	ip := strings.Split(r.RemoteAddr, ":")[0]

	_, exi := ser.BlackList[ip]
	if exi {
		http.Error(w, "403 forbidden", http.StatusForbidden)
		return false
	}

	ser.Lock.Lock()
	_, exi = ser.IPList[ip]
	if exi {
		ser.IPList[ip]++
	} else {
		ser.IPList[ip] = 1
	}
	ser.Lock.Unlock()

	return true
}

func (ser *Server) refresh() {
	for ; ; time.Sleep(time.Second) {
		if time.Since(ser.Timer) > refreshTime {
			ser.Timer = time.Now()

			for i, v := range ser.IPList {
				log.Println(i, "access:", v)
				if v > limitAccess {
					ser.BlackList[i] = true
				}

				delete(ser.IPList, i)
			}
		}
	}
}
