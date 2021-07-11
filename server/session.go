package server

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	// CookieName represents the name of cookie
	CookieName = "NTOU-SELL-ID"
)

var store = sessions.NewCookieStore([]byte(CookieName))

func login(w http.ResponseWriter, r *http.Request, uid int) {
	session, err := store.Get(r, CookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["auth"] = uid

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	}
}

func sessionValid(w http.ResponseWriter, r *http.Request) (int, bool) {
	session, err := store.Get(r, CookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return -1, false
	}

	auth := session.Values["auth"]

	if auth != nil {
		isAuth, ok := auth.(int)
		return isAuth, ok
	}

	http.Error(w, "請先登入!", http.StatusUnauthorized)
	return -1, false
}
