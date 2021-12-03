package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"se/server/backend"

	"github.com/gorilla/mux"
)

type leader struct {
	Name       string `json:"name"`
	SelfPoint  string `json:"self_point"`
	EnemyPoint string `json:"enemy_point"`
	Strength   string `json:"strength`
}

func fetchLeaderBoard(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {
		if mux.Vars(r)["key"] == "post" {
			decoder := json.NewDecoder(r.Body)
			var tmp leader
			err := decoder.Decode(&tmp)
			if err != nil {
				log.Println(err)
			}

			_, err = backend.AddLeader(tmp.Name, tmp.SelfPoint, tmp.EnemyPoint, tmp.Strength)
			if err == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, "ok")
			} else {
				http.Error(w, "failed", http.StatusBadRequest)
			}
		}
		return
	}

	path := mux.Vars(r)
	// args := r.URL.Query()

	switch path["key"] {
	case "get":
		leaders, err := backend.GetLeaders()
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
		} else {
			fmt.Fprint(w, leaders)
		}
	default:
		http.NotFound(w, r)
	}
}
