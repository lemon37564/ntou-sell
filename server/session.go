package server

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	IDLen      = 32
	CookieName = "sessID"
)

var store = sessions.NewCookieStore([]byte(CookieName))

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, CookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["auth"] = true
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("logged in")
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, CookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["auth"] = nil

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("logged out")
}

func sessionValid(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, CookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	auth := session.Values["auth"]
	if auth != nil {
		isAuth, ok := auth.(bool)
		return ok && isAuth
	}

	return false
}

func sessionHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "s1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(session)
	session.Values["id"] = genSessID()
	session.Save(r, w)
}

func genSessID() string {
	id := make([]byte, IDLen)

	if _, err := rand.Read(id); err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(id)
}
