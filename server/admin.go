package server

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/http"
	"se/database"

	"github.com/gorilla/mux"
)

const salt = "ntou-sell"

func sha512Hash(key string) string {
	key += salt

	hasher := sha512.New()
	hasher.Write([]byte(key))

	t := hasher.Sum(nil)
	return base64.URLEncoding.EncodeToString(t)
}

func adminDB(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)

	if r.Method == "POST" {
		if path["key"] == "exec" {
			query := r.FormValue("query")
			password := r.FormValue("admin-password")

			if sha512Hash(password) == "yqCijvyNB0TPJR45gSryOk2kgh3nYDCfSQovcZcnJ-36yFudgjCBLY4Rb-26Jyexd9F5vV9Ws93D1pd-cs-c7g==" {
				err := database.DirectAccess(query)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				} else {
					fmt.Fprint(w, "ok", err)
				}
			} else {
				http.Error(w, "failed", http.StatusUnauthorized)
			}
		}
		return
	}

	switch path["key"] {
	default:
		http.Error(w, "not found", http.StatusNotFound)
	}
}
