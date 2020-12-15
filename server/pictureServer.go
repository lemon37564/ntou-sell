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
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, PicHelp)
	case "upload":
		picName, exi := args["picname"]
		if exi {
			ser.picUpload(w, r, picName[0])

		} else {
			fmt.Fprint(w, "argument error")
		}

	case "show":
		picName, exi := args["picname"]
		if exi {
			ser.picShow(w, r, picName[0])

		} else {
			fmt.Fprint(w, "argument error")
		}
	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) picShow(w http.ResponseWriter, r *http.Request, picname string) {
	img, err := os.Open("webpage/img/" + picname)
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
	file, err := os.Create("webpage/img/" + picname)
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
