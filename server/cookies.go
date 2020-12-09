package server

import (
	"log"
	"net/http"
	"time"
)

func (ser *Server) setCookies(w http.ResponseWriter, r *http.Request, account, password string) {
	expire := time.Now()
	expire = expire.AddDate(1, 0, 0)

	cookie := &http.Cookie{Name: "seAccount", Value: account, Expires: expire, Secure: false}
	cookie2 := &http.Cookie{Name: "sePassword", Value: password, Expires: expire, Secure: false}

	http.SetCookie(w, cookie)
	r.AddCookie(cookie)

	http.SetCookie(w, cookie2)
	r.AddCookie(cookie2)

	redirectURL := "/"
	http.Redirect(w, r, redirectURL, 200)
}

func (ser *Server) getCookies(w http.ResponseWriter, r *http.Request) (*http.Cookie, *http.Cookie, bool) {
	cookie, err := r.Cookie("seAccount")
	cookie2, err2 := r.Cookie("sePassword")

	log.Println("cookie(account):", cookie, "err:", err)
	log.Println("cookie(password):", cookie2, "err:", err2)

	if err == nil && err2 == nil {
		return cookie, cookie2, true
	}

	return &http.Cookie{}, &http.Cookie{}, false
}
