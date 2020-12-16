package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func (ser *Server) picHandler(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	_, valid := sessionValid(w, r)
	if !valid {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)

	switch path["key"] {
	case "help":
		fmt.Fprint(w, PicHelp)
	case "upload":
		ser.picUpload(w, r, "test.jpg")
	case "get":
		ser.getPic(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) picUpload(w http.ResponseWriter, r *http.Request, picname string) {

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprint(w, handler.Header)
	f, err := os.Create("webpage/img/" + handler.Filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
}

func (ser *Server) getPic(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	psb := []string{".jpg", ".jpeg", ".png", ".webp", ".gif", ".ico", ".bmp"}

	pdid, exi := args["pdid"]
	if exi {
		_, err := strconv.Atoi(pdid[0])
		if err != nil {
			fmt.Fprint(w, "pdid is not an integer")
		}

		for _, v := range psb {
			_, err := os.Stat("webpage/img/" + pdid[0] + v)
			if err == nil {
				fmt.Fprint(w, "webpage/img/"+pdid[0]+v)
				return
			}
		}
	}

	fmt.Fprint(w, "not found")
}
