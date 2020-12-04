package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"se/user"
	"strings"
)

type server struct {
	db *sql.DB
	u  *user.User
}

func (ser *server) weber() {
	http.HandleFunc("/", ser.service)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (ser *server) service(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Printf("connection<host: %v, remote: %v>, receive: %v\n", r.Host, r.RemoteAddr, r.URL.Path)

	arg := strings.Split(r.URL.Path, "/")
	arg = arg[1:] // eliminate " "

	if len(arg) == 0 {
		http.NotFound(w, r)
		return
	}

	cmd := arg[0]

	if len(arg) == 1 && cmd == "all" {
		fmt.Fprintf(w, ser.u.GetAllUserData())
	} else if len(arg) == 2 && cmd == "login" {
		// format: /login/account=?&password=?
		acntpass := strings.Split(arg[1], "&")

		if len(acntpass) != 2 {
			fmt.Fprint(w, false)
		} else {
			acnt := strings.Split(acntpass[0], "=")
			pass := strings.Split(acntpass[1], "=")

			if acnt[0] == "account" && pass[0] == "password" && ser.u.Login(acnt[1], pass[1]) {
				fmt.Fprint(w, true)
			} else {
				fmt.Fprint(w, false)
			}
		}
	} else if len(arg) == 2 && cmd == "regist" {
		// format: /regist/account=?&password=?&name=?
		acntpass := strings.Split(arg[1], "&")

		if len(acntpass) != 3 {
			fmt.Fprint(w, "error")
		} else {
			acnt := strings.Split(acntpass[0], "=")
			pass := strings.Split(acntpass[1], "=")
			name := strings.Split(acntpass[2], "=")

			if acnt[0] == "account" && pass[0] == "password" && name[0] == "name" {
				fmt.Fprint(w, ser.u.Regist(acnt[1], pass[1], name[1]))
			} else {
				fmt.Fprint(w, "error")
			}
		}
	} else {
		http.NotFound(w, r)
	}
}
