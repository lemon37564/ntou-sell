package server

import (
	"fmt"
	"log"
	"net/http"
	"se/server/ai"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type request struct {
	board string
	color string
	level string
}

var server_ai *ai.AI8
var cache map[request]string = make(map[request]string)
var cacheLock sync.RWMutex

func ai_move(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)
	args := r.URL.Query()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if path["key"] == "move" {
		var level ai.Level
		start := time.Now()

		str_lv := args.Get("level")
		color := args.Get("color")
		board := args.Get("board")

		cacheLock.RLock()
		value, exist := cache[request{board, color, str_lv}]
		cacheLock.RUnlock()

		if exist {
			// requested but still computing
			if value == "" {
				http.Error(w, "not done yet", http.StatusBadRequest)
			} else {
				log.Println("cache hit, cache length:", len(cache))
				value += fmt.Sprintf(", time:%v", time.Since(start))
				fmt.Fprint(w, value)
			}
			return
		}

		if str_lv == "0" {
			level = ai.LV_ONE
		} else if str_lv == "1" {
			level = ai.LV_TWO
		} else {
			level = ai.LV_THREE
		}

		if color == "black" {
			server_ai = ai.NewAI8(ai.BLACK, level)
		} else {
			server_ai = ai.NewAI8(ai.WHITE, level)
		}

		cacheLock.Lock()
		cache[request{board, color, str_lv}] = ""
		cacheLock.Unlock()

		res, detail, err := server_ai.Move(board)
		if err != nil {
			fmt.Fprintf(w, "00, err:%v", err)
			return
		}
		x := res[0] - 'A'
		y := res[1] - 'a'

		duration := time.Since(start)
		if duration < time.Millisecond*300 {
			time.Sleep(time.Millisecond*300 - duration)
		}

		cacheLock.Lock()
		cache[request{board, color, str_lv}] = fmt.Sprintf("%d%d, value: %s", int(y), int(x), detail)
		cacheLock.Unlock()
		fmt.Fprintf(w, "%d%d, value: %s, time: %v", int(y), int(x), detail, duration)
	}

}
