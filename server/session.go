package server

import (
	"crypto/rand"
	"encoding/base64"
	"log"
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

type Session struct {
	lock sync.Mutex

	// use to record valid sessions
	list map[string]time.Time
}

func (se *Session) sessionValid(w http.ResponseWriter, r *http.Request) bool {
	se.lock.Lock()
	defer se.lock.Unlock()

	cookie, exist := getCookies(w, r)
	if exist {
		val, exi := se.list[cookie.Value]
		return exi && val.After(time.Now())
	}

	return false
}

func (se *Session) setSessionID(w http.ResponseWriter, r *http.Request) {
	se.lock.Lock()
	defer se.lock.Unlock()

	id := se.genSessID()
	setCookies(w, r, id)

	se.list[id] = time.Now().Add(lifeTime)
}

func (se *Session) genSessID() string {
	id := make([]byte, idLen)

	if _, err := rand.Read(id); err != nil {
		panic(err)
	}

	log.Println(base64.URLEncoding.EncodeToString(id))
	return base64.URLEncoding.EncodeToString(id)
}
