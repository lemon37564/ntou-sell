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

	path := r.URL.Path
	query := r.URL.Query()

	log.Printf("<host: %v, remote: %v>, path: %v, args: %v\n", r.Host, r.RemoteAddr, path, query)

	arg := strings.Split(path, "/")
	arg = arg[1:] // eliminate " "

	if len(arg) == 0 {
		http.NotFound(w, r)
		return
	}

	switch arg[0] {
	case "help":
		if len(arg) == 1 {
			fmt.Fprintln(w, `<html>
			<p> /user/all<br>
			列出所有帳號(僅限開發期間)<br>
			e.g.<br><a href=/user/all> 36.229.107.41/user/all </a><br><br>
			</p>
			<p> /user/login?account=...&password=...<br>
			登入是否成功(bool)<br>
			e.g.<br>36.229.107.41/login?account=test@gmail.com&password=000000<br><br>
			</p>
			<p> /user/regist?account=...&password=...&name=...<br>
			註冊新帳號<br>
			e.g.<br>36.229.107.41/regist?account=test2@gmail.com&password=1234&name=Wilson<br><br>
			<p> /user/delete?account=...&password=...<br>
			刪除帳號<br>
			e.g.<br>36.229.107.41/delete?account=test2@gmail.com&password=1234<br><br>
			</html>
			`)
		} else {
			http.NotFound(w, r)
		}
	case "user":
		if len(arg) == 2 {
			if arg[1] == "all" {
				fmt.Fprintf(w, ser.u.GetAllUserData())
			} else if arg[1] == "login" {
				val, exi := query["account"]
				val2, exi2 := query["password"]

				if exi && exi2 {
					fmt.Fprint(w, ser.u.Login(val[0], val2[0]))
				} else {
					fmt.Fprint(w, "argument error")
				}

			} else if arg[1] == "delete" {
				val, exi := query["account"]
				val2, exi2 := query["password"]

				if exi && exi2 {
					fmt.Fprint(w, ser.u.DeleteUser(val[0], val2[0]))
				} else {
					fmt.Fprint(w, "argument error")
				}

			} else if arg[1] == "regist" {
				val, exi := query["account"]
				val2, exi2 := query["password"]
				val3, exi3 := query["name"]

				if exi && exi2 && exi3 {
					fmt.Fprint(w, ser.u.Regist(val[0], val2[0], val3[0]))
				} else {
					fmt.Fprint(w, "argument error")
				}

			} else {
				http.NotFound(w, r)
			}
		} else {
			http.NotFound(w, r)
		}
	default:
		http.NotFound(w, r)
	}

}
