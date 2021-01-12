package server

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"se/database"
	"se/server/backend"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// Server handle all services
type Server struct {
	Ur *backend.User
	Pd *backend.Product
	Od *backend.Order
	Ht *backend.History
	Bd *backend.Bid
	Ct *backend.Cart
	Ms *backend.Message
}

// NewServer creates a new server
func NewServer() *Server {
	data := database.OpenAndInit()

	return &Server{
		Ur: backend.UserInit(data),
		Pd: backend.ProductInit(data),
		Od: backend.OrderInit(data),
		Ht: backend.HistoryInit(data),
		Bd: backend.BidInit(data),
		Ct: backend.CartInit(data),
		Ms: backend.MessageInit(data)}
}

// Serve start all functions provided for user
func (ser Server) Serve() {
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
	r.HandleFunc("/backend/user/{key}", ser.fetchUser)
	r.HandleFunc("/backend/message/{key}", ser.fetchMessage)

	r.HandleFunc("/backend/pics/{key}", ser.picHandler)

	http.Handle("/backend/", r)

	fs := http.FileServer(http.Dir("webpage"))
	http.Handle("/", ser.middleware(fs))

	go refresh()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (ser Server) middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !ser.validation(w, r) {
			return
		}

		h.ServeHTTP(w, r)
	})
}

// sizeof(void) = 0
type void struct{}

var (
	ipList   = make(map[string]int)
	blockSet = make(map[string]void)

	// read write mutex (prevent race condition)
	// it is forbidden to read and write the map concurrently
	ipLock    sync.RWMutex
	blockLock sync.RWMutex
)

func (ser Server) validation(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()

	ip := getIP(r)

	blockLock.RLock()
	_, exi := blockSet[ip]
	blockLock.RUnlock()
	if exi {
		http.Error(w, "403 forbidden", http.StatusForbidden)
		return false
	}

	// read lock
	ipLock.RLock()
	_, exi = ipList[ip]
	ipLock.RUnlock()

	// write lock
	ipLock.Lock()
	if exi {
		ipList[ip]++
	} else {
		ipList[ip] = 1
	}
	ipLock.Unlock()

	return true
}

// a thread refreshing the blockSet and ipList
func refresh() {
	const (
		limitAccess = 150
		refreshTime = time.Second * 30
	)

	for loop := 0; ; time.Sleep(refreshTime) {

		// write lock
		blockLock.Lock()
		// unban (30min)
		if loop%60 == 0 {
			// leave the old one to GC
			blockSet = make(map[string]void)
		}
		blockLock.Unlock()

		// read lock
		ipLock.RLock()
		for i, v := range ipList {
			if v > limitAccess {
				// write lock
				blockLock.Lock()
				blockSet[i] = void{}
				blockLock.Unlock()

				log.Printf("ip %15s access %5d times in 30s, banned.\n", i, v)
			} else {
				log.Printf("ip %15s access %5d times in 30s\n", i, v)
			}
		}
		ipLock.RUnlock()

		// leave the old one to GC
		ipList = make(map[string]int)
		loop++
	}
}

func getIP(r *http.Request) string {
	ipAdress := r.Header.Get("X-Real-Ip")
	if ipAdress == "" {
		ipAdress = r.Header.Get("X-Forwarded-For")
	}
	if ipAdress == "" {
		ipAdress = r.RemoteAddr
	}
	return ipAdress
}
