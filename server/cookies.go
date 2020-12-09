package server

import (
	"net/http"
	"time"
)

func (ser *Server) setCookies(w http.ResponseWriter, r *http.Request, account, password string) {
	expire := time.Now()
	expire = expire.AddDate(365, 0, 0)

	cookie := http.Cookie{Name: "account", Value: account, Expires: expire}
	cookie2 := http.Cookie{Name: "password", Value: password, Expires: expire}

	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}

func (ser *Server) getCookies(w http.ResponseWriter, r *http.Request) (*http.Cookie, *http.Cookie, bool) {
	cookie, err := r.Cookie("account")
	cookie2, err2 := r.Cookie("password")

	if err == nil && err2 == nil {
		return cookie, cookie2, true
	}

	return &http.Cookie{}, &http.Cookie{}, false
}
