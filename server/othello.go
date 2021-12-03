package server

import (
	"fmt"
	"net/http"
	"se/server/backend"

	"github.com/gorilla/mux"
)

func fetchLeaderBoard(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if mux.Vars(r)["key"] == "post" {
			r.ParseMultipartForm(32 << 20)

			name := r.FormValue("name")
			point := r.FormValue("point")
			strength := r.FormValue("strength")

			_, err := backend.AddLeader(name, point, strength)
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
