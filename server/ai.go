package server

import (
	"fmt"
	"net/http"
	"se/server/ai"
	"time"

	"github.com/gorilla/mux"
)

var server_ai *ai.AI8

func ai_move(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)
	args := r.URL.Query()

	if path["key"] == "move" {
		var level ai.Level
		str_lv := args.Get("level")
		if str_lv == "0" {
			level = ai.LV_ONE
		} else if str_lv == "1" {
			level = ai.LV_TWO
		} else {
			level = ai.LV_THREE
		}

		color := args.Get("color")
		if color == "black" {
			server_ai = ai.NewAI8(ai.BLACK, level)
		} else {
			server_ai = ai.NewAI8(ai.WHITE, level)
		}

		board := args.Get("board")
		fmt.Println(str_lv, color, board)

		start := time.Now()

		res, detail, err := server_ai.Move(board)
		if err != nil {
			fmt.Fprintf(w, "00, err:%v", err)
			return
		}
		x := res[0] - 'A'
		y := res[1] - 'a'

		duration := time.Since(start)
		if duration < time.Millisecond*200 {
			time.Sleep(time.Millisecond*200 - duration)
		}

		fmt.Fprintf(w, "%d%d, value: %s, time: %v", int(y), int(x), detail, duration)
	}

}
