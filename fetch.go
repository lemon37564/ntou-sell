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
		fmt.Fprintf(w, ser.us.GetAllUserData())
	case "login":
		val, exi := args["account"]
		val2, exi2 := args["password"]

		if exi && exi2 {
			fmt.Fprint(w, ser.us.Login(val[0], val2[0]))
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "delete":
		val, exi := args["account"]
		val2, exi2 := args["password"]

		if exi && exi2 {
			fmt.Fprint(w, ser.us.DeleteUser(val[0], val2[0]))
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "regist":
		val, exi := args["account"]
		val2, exi2 := args["password"]
		val3, exi3 := args["name"]

		if exi && exi2 && exi3 {
			fmt.Fprint(w, ser.us.Regist(val[0], val2[0], val3[0]))
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
		fmt.Fprint(w, ser.pr.GetAll())
	case "add":
		exist := make([]bool, 7)
		var name, price, des, amount, account, bid, date []string

		name, exist[0] = args["name"]
		price, exist[1] = args["price"]
		des, exist[2] = args["description"]
		amount, exist[3] = args["amount"]
		account, exist[4] = args["account"]
		bid, exist[5] = args["bid"]
		date, exist[6] = args["date"]

		if all(exist) {
			p, err1 := strconv.Atoi(price[0])
			a, err2 := strconv.Atoi(amount[0])
			b := (bid[0] == "true")

			if err1 == nil && err2 == nil {
				fmt.Fprint(w, ser.pr.AddProduct(name[0], p, des[0], a, account[0], b, date[0]))
			} else {
				fmt.Fprint(w, "price or amount was not an integer.")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "delete":
	case "search":
		val, exi := args["name"]

		if exi {
			fmt.Fprint(w, ser.pr.SearchProductsByName(val[0]))
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "filter_search":
		exist := make([]bool, 4)
		var name, min, max, eval []string

		name, exist[0] = args["name"]
		min, exist[1] = args["minprice"]
		max, exist[2] = args["maxprice"]
		eval, exist[3] = args["eval"]

		if all(exist) {
			mi, err1 := strconv.Atoi(min[0])
			ma, err2 := strconv.Atoi(max[0])
			ev, err3 := strconv.Atoi(eval[0])

			if err1 == nil && err2 == nil && err3 == nil {
				fmt.Fprint(w, ser.pr.EnhanceSearchProductsByName(name[0], mi, ma, ev))
			} else {
				fmt.Fprint(w, "min price, max price or evaluation was not as interger")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
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
				<a href=/user/all> /user/all </a><br><br>
			</p>
			<p> 
				/user/login?account=...&password=...<br>
				登入是否成功(bool)<br>
				e.g.<br><a href=https://se-ssb.herokuapp.com/user/login?account=test@gmail.com&password=0000>
				https://se-ssb.herokuapp.com/user/login?account=test@gmail.com&password=0000</a>
				<br><br>
			</p>
			<p>
				/user/regist?account=...&password=...&name=...<br>
				註冊新帳號<br>
				e.g.<br><a href=https://se-ssb.herokuapp.com/user/regist?account=test2@gmail.com&password=1234&name=Wilson>
				https://se-ssb.herokuapp.com/user/regist?account=test2@gmail.com&password=1234&name=Wilson</a>
				<br><br>
			</p>
			<p>
				/user/delete?account=...&password=...<br>
				刪除帳號<br>
				e.g.<br><a href=https://se-ssb.herokuapp.com/user/delete?account=test3@gmail.com&password=1234>
				https://se-ssb.herokuapp.com/user/delete?account=test3@gmail.com&password=1234</a>
				<br><br>
			</p>
			<p> 
				/product/all<br>
				列出所有商品(僅限開發期間)<br>
				<a href=/product/all> /product/all </a><br><br>
			</p>
			<p> 
				/product/add?name=...&price=...&description=...&amount=...&account=...&bid=...&date=...<br>
				新增商品<br>
				e.g.<br><a href=https://se-ssb.herokuapp.com/product/add?name=ifone12&price=5000&description=盜版商品&amount=10&account=test@gmail.com&bid=true&date=2020-12-31>
				https://se-ssb.herokuapp.com/product/add?name=ifone12&price=5000&description=盜版商品&amount=10&account=test@gmail.com&bid=true&date=2020-12-31</a>
				<br><br>
			</p>
			<p>
				/product/search?name=...<br>
				查詢商品<br>
				e.g.<br><a href=https://se-ssb.herokuapp.com/product/search?name=ifone>
				https://se-ssb.herokuapp.com/product/search?name=ifone</a>
				<br><br>
			</p>
			<p>
				/product/search?name=...&minprice=...&maxprice=...&eval=...<br>
				查詢商品(過濾)<br>
				e.g.<br><a href=https://se-ssb.herokuapp.com/product/search?name=ifone&minprice=10&maxprice=5000&eval=2>
				https://se-ssb.herokuapp.com/product/search?name=ifone&minprice=10&maxprice=5000&eval=2</a>
				<br><br>
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
