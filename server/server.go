package server

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

func Serve() {
	port := os.Getenv("PORT")

	// when test on localhost
	if port == "" {
		port = "8080"
	}
	log.Println("Service running on port:", port)

	r := mux.NewRouter()

	r.HandleFunc("/backend/{key}", defaultFunc)
	r.HandleFunc("/backend/bid/{key}", fetchBid)
	r.HandleFunc("/backend/cart/{key}", fetchCart)
	r.HandleFunc("/backend/history/{key}", fetchHistory)
	r.HandleFunc("/backend/order/{key}", fetchOrder)
	r.HandleFunc("/backend/product/{key}", fetchProduct)
	r.HandleFunc("/backend/user/{key}", fetchUser)
	r.HandleFunc("/backend/message/{key}", fetchMessage)
	r.HandleFunc("/backend/ai/{key}", ai_move)
	r.HandleFunc("/backend/leaderboard/{key}", fetchLeaderBoard)
	r.HandleFunc("/backend/db/{key}", adminDB)

	r.HandleFunc("/backend/pics/{key}", picHandler)

	http.Handle("/backend/", r)

	fs := http.FileServer(http.Dir("webpage"))
	http.Handle("/", middleware(fs))

	go refresh()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !validation(w, r) {
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

func validation(w http.ResponseWriter, r *http.Request) bool {
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
