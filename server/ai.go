package server

import (
	"fmt"
	"net/http"
	"se/server/ai"

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
			level = ai.LV_THREE
		} else {
			level = ai.LV_FIVE
		}

		color := args.Get("color")
		if color == "black" {
			server_ai = ai.NewAI8(ai.BLACK, level)
		} else {
			server_ai = ai.NewAI8(ai.WHITE, level)
		}

		board := args.Get("board")
		fmt.Println(str_lv, color, board)

		res, err := server_ai.Move(board)
		if err != nil {
			fmt.Println(err)
			return
		}
		x := res[0] - 'A'
		y := res[1] - 'a'
		fmt.Fprint(w, int(x), int(y))
	}

}
