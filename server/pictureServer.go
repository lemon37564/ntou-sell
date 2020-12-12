package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func (ser *Server) picHandler(w http.ResponseWriter, r *http.Request) {
	if !ser.Sess.sessionValid(w, r) {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	picPath := path["key"]

	img, err := os.Open("pics/" + picPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer img.Close()

	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, img)
}
