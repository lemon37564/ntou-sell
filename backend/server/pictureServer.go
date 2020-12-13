package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func (ser *Server) picHandler(w http.ResponseWriter, r *http.Request) {
	if !sessionValid(w, r) {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)

	switch path["key"] {
	case "upload":
		ser.picUpload(w, r, "temp.jpg")
	case "show":
		ser.picShow(w, r, "temp.jpg")
	default:
		http.NotFound(w, r)
	}

}

func (ser *Server) picShow(w http.ResponseWriter, r *http.Request, picname string) {
	img, err := os.Open("pics/" + picname)
	if err != nil {
		log.Println(err)
		return
	}
	defer img.Close()

	imgType := strings.Split(img.Name(), ".")[1]

	w.Header().Set("Content-Type", "image/"+imgType)
	_, err = io.Copy(w, img)
	if err != nil {
		log.Println(err)
	}
}

func (ser *Server) picUpload(w http.ResponseWriter, r *http.Request, picname string) {
	file, err := os.Create("pics/" + picname)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprint(w, "ok")
}
