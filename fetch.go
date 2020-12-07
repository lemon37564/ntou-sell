package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (ser *server) fetch(w http.ResponseWriter, r *http.Request, cmd string, args map[string][]string) {
	path := strings.Split(cmd, "/")

	if len(path) == 0 {
		http.NotFound(w, r)
	}

	switch path[0] {
	case "help":
		if len(path) == 1 {
			help(w)
		} else {
			http.NotFound(w, r)
		}
	case "user":
		ser.fetchUser(w, r, path, args)
	case "product":
		ser.fetchProduct(w, r, path, args)
	default:
		http.NotFound(w, r)
	}
}

func (ser *server) fetchUser(w http.ResponseWriter, r *http.Request, path []string, args map[string][]string) {
	if len(path) != 2 {
		http.NotFound(w, r)
		return
	}

	switch path[1] {
	case "all":
		fmt.Fprintf(w, ser.u.GetAllUserData())
	case "login":
		val, exi := args["account"]
		val2, exi2 := args["password"]

		if exi && exi2 {
			fmt.Fprint(w, ser.u.Login(val[0], val2[0]))
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "delete":
		val, exi := args["account"]
		val2, exi2 := args["password"]

		if exi && exi2 {
			fmt.Fprint(w, ser.u.DeleteUser(val[0], val2[0]))
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "regist":
		val, exi := args["account"]
		val2, exi2 := args["password"]
		val3, exi3 := args["name"]

		if exi && exi2 && exi3 {
			fmt.Fprint(w, ser.u.Regist(val[0], val2[0], val3[0]))
		} else {
			fmt.Fprint(w, "argument error")
		}
	default:
		http.NotFound(w, r)
	}
}

func (ser *server) fetchProduct(w http.ResponseWriter, r *http.Request, path []string, args map[string][]string) {
	if len(path) != 2 {
		http.NotFound(w, r)
		return
	}

	switch path[1] {
	case "all":
	case "add":
	case "delete":
	case "search":
		val, exi := args["name"]

		if exi {
			_ = val
			// fmt.Fprint(w, ser.p.)
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "search-filter":
	default:
		http.NotFound(w, r)
	}
}

func help(w http.ResponseWriter) {
	fmt.Fprintln(w, `
		<html>
			<p> 
				/user/all<br>
				列出所有帳號(僅限開發期間)<br>
				e.g.<br><a href=/user/all> 36.229.107.41/user/all </a><br><br>
			</p>
			<p> 
				/user/login?account=...&password=...<br>
				登入是否成功(bool)<br>
				e.g.<br>36.229.107.41/login?account=test@gmail.com&password=000000<br><br>
			</p>
			<p>
				/user/regist?account=...&password=...&name=...<br>
				註冊新帳號<br>
				e.g.<br>36.229.107.41/regist?account=test2@gmail.com&password=1234&name=Wilson<br><br>
			</p>
			<p>
				/user/delete?account=...&password=...<br>
				刪除帳號<br>
				e.g.<br>36.229.107.41/delete?account=test2@gmail.com&password=1234<br><br>
			</p>
		</html>
		`)
}
