package server

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"se/server/backend"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var authKeys = make(map[string]time.Time)
var authLock sync.RWMutex

func fetchLeaderBoard(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "genKey":
		key := make([]byte, 32)
		rand.Read(key)
		keyStr := base64.URLEncoding.EncodeToString(key)

		fmt.Fprint(w, keyStr)

		const secret = "wp1101-final-0076D053-00771053"
		hashed := sha256.Sum256([]byte(keyStr + secret))
		authLock.Lock()
		cleanExpiredKeys()
		authKeys[hex.EncodeToString(hashed[:])] = time.Now()
		authLock.Unlock()
	case "get":
		leaders, err := backend.GetLeaders()
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
		} else {
			fmt.Fprint(w, leaders)
		}
	case "add":
		name := args.Get("name")
		selfPoint := args.Get("self_point")
		enemyPoint := args.Get("enemy_point")
		strength := args.Get("strength")
		verification := args.Get("verification")

		authLock.Lock()
		cleanExpiredKeys()
		_, exist := authKeys[verification]
		if exist {
			delete(authKeys, verification)
		}
		authLock.Unlock()

		if !exist {
			log.Println("verification not accepted, score disposed")
			http.Error(w, "verfication failed", http.StatusNotAcceptable)
			return
		}
		log.Println("verification accepted, score have been uploaded")

		_, err := backend.AddLeader(name, selfPoint, enemyPoint, strength)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "ok")
		} else {
			http.Error(w, "failed", http.StatusBadRequest)
		}
	default:
		http.NotFound(w, r)
	}
}

func cleanExpiredKeys() {
	for key, t := range authKeys {
		if time.Since(t) > time.Second*5 {
			delete(authKeys, key)
		}
	}
}
