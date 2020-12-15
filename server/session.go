package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/sessions"
)

const (
	CookieName  = "SESSID"
	limitAccess = 30
	refreshTime = time.Second * 10
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
		return
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
		return
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

	return -1, false
}

func (ser *Server) validation(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()

	ip := strings.Split(r.RemoteAddr, ":")[0]

	_, exi := ser.BlackList[ip]
	if exi {
		http.Error(w, "BLOCKED", http.StatusForbidden)
		return false
	}

	_, exi = ser.IPList[ip]
	if exi {
		ser.IPList[ip]++
	} else {
		ser.IPList[ip] = 1
	}

	if time.Since(ser.Timer) > refreshTime {
		ser.Timer = time.Now()

		for i, v := range ser.IPList {
			if v > limitAccess {
				ser.BlackList[ip] = true
			}

			delete(ser.IPList, i)
		}
	}

	return true
}
