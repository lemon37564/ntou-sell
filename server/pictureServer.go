package server

import (
	"io"
	"log"
	"net/http"
	"os"
)

func picHandler(w http.ResponseWriter, r *http.Request) {
	img, err := os.Open("pics/test.png")
	if err != nil {
		log.Println(err)
		return
	}
	defer img.Close()

	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, img)
}
