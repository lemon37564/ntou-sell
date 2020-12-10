package server

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
	"time"
)

const (
	idLen       = 32
	cookieName  = "sessID"
	lifeTime    = time.Hour * 24
	refreshTime = time.Hour
)

type session struct {
	lock sync.Mutex

	// use to record valid sessions
	list map[string]time.Time

	lastRefresh time.Time
}

// NewSession return a session handler
func NewSession() *session {
	s := new(session)
	s.list = make(map[string]time.Time)

	return s
}

func (se *session) sessionValid(w http.ResponseWriter, r *http.Request) bool {
	se.lock.Lock()
	defer se.lock.Unlock()

	cookie, exist := getCookies(w, r)
	if exist {
		_, exi := se.list[cookie.Value]
		return exi
	}

	return false
}

func (se *session) setSessionID(w http.ResponseWriter, r *http.Request) {
	se.lock.Lock()
	defer se.lock.Unlock()

	id := se.genSessID()
	setCookies(w, r, id)

	se.list[id] = time.Now().Add(lifeTime)
	http.Redirect(w, r, `/success`, 301)
}

func (se *session) sessionRefresh() {
	// check sessions is valid (delete it if not)
	if time.Since(se.lastRefresh) > refreshTime {
		now := time.Now()
		se.lastRefresh = now

		for i, v := range se.list {
			if now.After(v) {
				delete(se.list, i)
			}
		}
	}
}

func (se *session) genSessID() string {
	id := make([]byte, idLen)

	if _, err := rand.Read(id); err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(id)
}
