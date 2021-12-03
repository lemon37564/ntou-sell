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
	Name     string `json:"name"`
	Point    string `json:"point"`
	Strength string `json:"strength`
}

func fetchLeaderBoard(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if mux.Vars(r)["key"] == "post" {

			decoder := json.NewDecoder(r.Body)
			var tmp leader
			err := decoder.Decode(&tmp)
			if err != nil {
				log.Println(err)
			}

			_, err = backend.AddLeader(tmp.Name, tmp.Point, tmp.Strength)
			if err == nil {
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
