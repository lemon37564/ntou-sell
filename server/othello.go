package server

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"se/server/backend"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var authKeys = make(map[string]authData)
var authLock sync.RWMutex

type player struct {
	Name       string `json:"name"`
	SelfPoint  string `json:"self_point"`
	EnemyPoint string `json:"enemy_point"`
	Strength   string `json:"strength"`
}

type authData struct {
	expiredTime time.Time
	selfPoint   string
	enemyPoint  string
	strength    string
}

func fetchLeaderBoard(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "genKey":
		key := make([]byte, 32)
		rand.Read(key)
		keyStr := base64.URLEncoding.EncodeToString(key)

		selfPoint := args.Get("self")
		enemyPoint := args.Get("enemy")
		strength := args.Get("str")

		fmt.Fprint(w, keyStr)

		secret := os.Getenv("SECRET_SALT_KEY")
		hashed := sha256.Sum256([]byte(keyStr + secret + ";" + selfPoint + ";" + enemyPoint + ";" + strength))
		authLock.Lock()
		cleanExpiredKeys()
		authKeys[hex.EncodeToString(hashed[:])] = authData{
			expiredTime: time.Now().Add(time.Second * 5),
			selfPoint:   selfPoint,
			enemyPoint:  enemyPoint,
			strength:    strength,
		}
		authLock.Unlock()
	case "getRaw":
		leaders, err := backend.GetLeadersRaw()
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
		} else {
			fmt.Fprint(w, leaders)
		}
	case "get":
		strength := args.Get("strength")
		amount := args.Get("amount")
		t := time.Now()
		leaders, err := backend.GetLeadersOrdered(strength, amount)
		if err != nil {
			log.Println(err)
			http.Error(w, "error", http.StatusInternalServerError)
		} else {
			fmt.Fprint(w, leaders)
			log.Printf("query spent: %v\n", time.Since(t))
		}
	case "add":
		if r.Header.Get("referer") != "https://lemon37564.github.io/" {
			log.Println("referer not accepted:", r.Header.Get("referer"))
			http.Error(w, "error", http.StatusForbidden)
			return
		}

		value := args.Get("v")
		verification := args.Get("verification")

		// 123-45-31-90... -> ["123", "45", "31", "90", ...]
		valueDecode, err := url.QueryUnescape(value)
		if err != nil {
			log.Println("url decode failed")
			log.Println("format not fit")
			http.Error(w, "format not fit", http.StatusBadRequest)
			return
		}

		p := player{}
		err = json.Unmarshal([]byte(valueDecode), &p)
		if err != nil {
			http.Error(w, "failed", http.StatusBadRequest)
			return
		}
		log.Println(p)

		var data authData
		authLock.Lock()
		// if key timeout, delete
		cleanExpiredKeys()
		data, exist := authKeys[verification]
		if exist {
			delete(authKeys, verification)
		}
		authLock.Unlock()

		// if not exist or if didn't match the data(prevent packet sniffering)
		if !exist ||
			data.selfPoint != p.SelfPoint ||
			data.enemyPoint != p.EnemyPoint ||
			data.strength != p.Strength {
			log.Println("verification not accepted, score disposed")
			http.Error(w, "verfication failed", http.StatusNotAcceptable)
			return
		}
		log.Println("verification accepted, score have been uploaded")

		_, err = backend.AddLeader(p.Name, p.SelfPoint, p.EnemyPoint, p.Strength)
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
	for key, dat := range authKeys {
		if time.Now().After(dat.expiredTime) {
			delete(authKeys, key)
		}
	}
}
