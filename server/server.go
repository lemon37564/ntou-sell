package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"runtime"
	"se/server/backend"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

const (
	limitAccess = 90
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
}

type void struct {
	// sizeof(void) = 0
}

var (
	IPList   = make(map[string]int)
	BlockSet = make(map[string]void)
	Lock     sync.Mutex
)

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

	fs := http.FileServer(http.Dir("webpage"))
	http.Handle("/", ser.middleware(fs))

	go ser.refresh()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (ser *Server) middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !ser.validation(w, r) {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (ser *Server) validation(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()

	ip := ser.getIP(r)

	// has a possbility to panic: map concurrent read and write
	// when it is refreshing
	_, exi := BlockSet[ip]
	if exi {
		http.Error(w, "403 forbidden", http.StatusForbidden)
		return false
	}

	// mutex (prevent race condition)
	// it is forbidden to concurrent read and write
	Lock.Lock()

	if _, exi = IPList[ip]; exi {
		IPList[ip]++
	} else {
		IPList[ip] = 1
	}
	Lock.Unlock()

	return true
}

func (ser *Server) refresh() {
	for loop := 0; ; time.Sleep(refreshTime) {

		// unban(3min)
		if loop%6 == 0 {
			// leave the old one to GC
			BlockSet = make(map[string]void)
		}

		for i, v := range IPList {
			if v > limitAccess {
				BlockSet[i] = void{}
				log.Printf("ip %15s access %5d times in 30s, banned.\n", i, v)
			} else {
				log.Printf("ip %15s access %5d times in 30s\n", i, v)
			}
		}

		// leave the old one to GC
		IPList = make(map[string]int)
		loop++
	}
}

func (ser *Server) getIP(r *http.Request) string {
	ipAdress := r.Header.Get("X-Real-Ip")
	if ipAdress == "" {
		ipAdress = r.Header.Get("X-Forwarded-For")
	}
	if ipAdress == "" {
		ipAdress = r.RemoteAddr
	}
	return ipAdress
}
