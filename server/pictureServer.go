package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func (ser *Server) picHandler(w http.ResponseWriter, r *http.Request) {
	_, valid := sessionValid(w, r)
	if !valid {
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
	img, err := os.Open("server/backend/pics/" + picname)
	if err != nil {
		log.Println("opening file:", err)
		return
	}
	defer img.Close()

	imgType := strings.Split(img.Name(), ".")[1]

	w.Header().Set("Content-Type", "image/"+imgType)
	_, err = io.Copy(w, img)
	if err != nil {
		log.Println("showing file:", err)
	}
}

func (ser *Server) picUpload(w http.ResponseWriter, r *http.Request, picname string) {
	file, err := os.Create("server/backend/pics/" + picname)
	if err != nil {
		log.Println("creating file:", err)
		return
	}
	defer file.Close()

	v, _ := json.Marshal(r.Body)
	log.Println("request body:", v)

	_, err = io.Copy(file, r.Body)
	if err != nil {
		log.Println("copying file:", err)
		return
	}

	fmt.Fprint(w, "ok")
}
