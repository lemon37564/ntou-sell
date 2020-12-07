package main

import (
	"fmt"
	"net/http"
	"strconv"
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
		fmt.Fprint(w, ser.p.GetAll())
	case "add":
		exist := make([]bool, 7)
		var name, price, des, amount, account, bid, date []string
		var _p, _a int
		var _b bool

		name, exist[0] = args["name"]
		price, exist[1] = args["price"]
		des, exist[2] = args["description"]
		amount, exist[3] = args["amount"]
		account, exist[4] = args["account"]
		bid, exist[5] = args["bid"]
		date, exist[6] = args["date"]

		if all(exist) {
			_p, _ = strconv.Atoi(price[0])
			_a, _ = strconv.Atoi(amount[0])
			_b = (bid[0] == "true")
			fmt.Fprint(w, ser.p.AddProduct(name[0], _p, des[0], _a, account[0], _b, date[0]))
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "delete":
	case "search":
		val, exi := args["name"]

		if exi {
			fmt.Fprint(w, ser.p.SearchProductsByName(val[0]))
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
				e.g.<br><a href=/user/all> /user/all </a><br><br>
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
			<p> 
				/product/all<br>
				列出所有商品(僅限開發期間)<br>
				e.g.<br><a href=/product/all> /product/all </a><br><br>
			</p>
			<p> 
				/product/add?name=...&price=...&description=...&amount=...&account=...&bid=...&date=...<br>
				新增商品(bool)<br>
			</p>
			<p>
				/product/search?name=...<br>
				查詢商品<br>
			</p>
		</html>
		`)
}

func all(bs []bool) bool {
	for _, v := range bs {
		if !v {
			return false
		}
	}

	return true
}
