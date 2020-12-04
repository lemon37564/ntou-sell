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

	class := arg[0]
	length := len(arg)

	switch class {
	case "help":
		if length == 1 {
			fmt.Fprintln(w, `<html>
			<p> /user/all<br>
			列出所有帳號(僅限開發期間)<br>
			e.g.<br><a href=/user/all> 36.229.107.41/user/all </a><br><br>
			</p>
			<p> /user/login/account=?&password=?<br>
			登入是否成功(bool)<br>
			e.g.<br>36.229.107.41/login/account=test@gmail.com&password=000000<br><br>
			</p>
			<p> /user/regist/account=?&password=?&name=?<br>
			註冊新帳號<br>
			e.g.<br>36.229.107.41/regist/account=test2@gmail.com&password=1234&name=Wilson<br><br>
			<p> /user/delete/account=?&password=?<br>
			刪除帳號<br>
			e.g.<br>36.229.107.41/delete/account=test2@gmail.com&password=1234<br><br>
			</html>
			`)
		} else {
			http.NotFound(w, r)
		}
	case "user":
		if length == 2 && arg[1] == "all" {
			fmt.Fprintf(w, ser.u.GetAllUserData())
		} else if length == 3 {
			if arg[1] == "login" {
				acntpass := strings.Split(arg[2], "&")

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
			} else if arg[1] == "delete" {
				acntpass := strings.Split(arg[2], "&")

				if len(acntpass) != 2 {
					fmt.Fprint(w, "error")
				} else {
					acnt := strings.Split(acntpass[0], "=")
					pass := strings.Split(acntpass[1], "=")

					if acnt[0] == "account" && pass[0] == "password" {
						fmt.Fprint(w, ser.u.DeleteUser(acnt[1], pass[1]))
					} else {
						fmt.Fprint(w, "error")
					}
				}
			} else if arg[1] == "regist" {
				acntpass := strings.Split(arg[2], "&")

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
		} else {
			http.NotFound(w, r)
		}
	default:
		http.NotFound(w, r)
	}

}
