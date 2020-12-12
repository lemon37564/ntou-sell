package server

import (
	"log"
	"net/http"
	"time"
)

func setCookies(w http.ResponseWriter, r *http.Request, sid string) {
	expire := time.Now()
	expire = expire.AddDate(1, 0, 0)

	cookie := &http.Cookie{Name: cookieName, Value: sid, Expires: expire, Secure: false, Path: "/", Domain: "se-ssb.herokuapp.com"}

	http.SetCookie(w, cookie)
	r.AddCookie(cookie)

	w.WriteHeader(http.StatusOK)
}

func getCookies(w http.ResponseWriter, r *http.Request) (*http.Cookie, bool) {
	cookie, err := r.Cookie(cookieName)

	log.Println("cookie:", cookie, "err:", err)

	return cookie, err == nil
}
