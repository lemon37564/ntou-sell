package server

import (
	"fmt"
	"net/http"
	"se/server/backend"

	"github.com/gorilla/mux"
)

func fetchLeaderBoard(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
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
