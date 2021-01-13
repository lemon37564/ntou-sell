package server

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (ser Server) picHandler(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	_, valid := sessionValid(w, r)
	if !valid {
		return
	}

	path := mux.Vars(r)

	switch path["key"] {
	case "help":
		fmt.Fprint(w, PicAPI)
	case "upload":
		ser.picUpload(w, r)
	case "get":
		ser.getPic(w, r)
	case "changeBg":
		ser.changeBg(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (ser Server) picUpload(w http.ResponseWriter, r *http.Request) {

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

func (ser Server) getPic(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	// bad way, rewrite it later
	psb := []string{".jpg", ".jpeg", ".png", ".webp", ".gif", ".ico", ".bmp"}

	pdid := args.Get("pdid")
	_, err := strconv.Atoi(pdid)
	if err != nil {
		fmt.Fprint(w, "pdid is not an integer")
		return
	}

	for _, v := range psb {
		_, err := os.Stat("webpage/img/" + pdid + v)
		if err == nil {
			fmt.Fprint(w, pdid+v)
			return
		}
	}

	fmt.Fprint(w, "none.webp")
	log.Println("img not found")
}

func (ser Server) changeBg(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	f, err := os.Create("webpage/img/temp.webp")
	if err != nil {
		log.Println(err)
		return
	}
	io.Copy(f, file)
	f.Close()

	timeForm := r.FormValue("time")
	t, err := time.Parse("2012-12-31 23:59:59", timeForm)
	fmt.Fprint(w, t, err)

	go func(t time.Time, file multipart.File) {
		for ; ; time.Sleep(time.Second) {
			if time.Now().After(t) {
				os.Remove("webpage/img/bg2.webp")
				os.Rename("webpage/img/temp.webp", "webpage/img/bg2.webp")
				return
			}
		}
	}(t, file)

	fmt.Fprint(w, "ok")
}
