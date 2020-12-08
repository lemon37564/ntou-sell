package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (ser *server) fetch(cmd string, args map[string][]string) {
	path := strings.Split(cmd, "/")

	if len(path) == 0 {
		http.NotFound(ser.w, ser.r)
	}

	switch path[0] {
	case "help":
		if len(path) == 1 {
			ser.help()
		} else {
			http.NotFound(ser.w, ser.r)
		}
	case "user":
		ser.fetchUser(path, args)
	case "product":
		ser.fetchProduct(path, args)
	default:
		http.NotFound(ser.w, ser.r)
	}
}

func (ser *server) fetchHistory(path []string, args map[string][]string) {
	if len(path) != 2 {
		http.NotFound(ser.w, ser.r)
		return
	}

	switch path[1] {
	case "all":
		fmt.Fprint(ser.hi.GetAll())
	case "get":
		ac, exi := args["account"]
		val, exist := args["amount"]

		if exist && exi {
			amnt, err := strconv.Atoi(val[0])
			if err == nil {
				uid := ser.us.GetUIDFromAccount(ac[0])
				fmt.Fprint(ser.w, ser.hi.GetHistory(uid, amnt))
			} else {
				fmt.Fprint(ser.w, "amount was not an integer")
			}
		} else {
			fmt.Fprint(ser.w, "argument error")
		}
	case "delete":
		ac, exi := args["account"]
		pdid, exi2 := args["pd_id"]

		if exi && exi2 {
			pd, err := strconv.Atoi(pdid[0])
			if err == nil {
				uid := ser.us.GetUIDFromAccount(ac[0])
				fmt.Fprint(ser.w, ser.hi.GetHistory(uid, pd))
			} else {
				fmt.Fprint(ser.w, "pd_id was not an integer")
			}
		} else {
			fmt.Fprint(ser.w, "argument error")
		}
	case "add":
		ac, exi := args["account"]
		pdid, exi2 := args["pd_id"]

		if exi && exi2 {
			pd, err := strconv.Atoi(pdid[0])
			if err == nil {
				uid := ser.us.GetUIDFromAccount(ac[0])
				fmt.Fprint(ser.w, ser.hi.AddHistory(uid, pd))
			} else {
				fmt.Fprint(ser.w, "pd_id was not an integer")
			}
		} else {
			fmt.Fprint(ser.w, "argument error")
		}
	default:
		http.NotFound(ser.w, ser.r)
	}
}

func (ser *server) fetchUser(path []string, args map[string][]string) {
	if len(path) != 2 {
		http.NotFound(ser.w, ser.r)
		return
	}

	switch path[1] {
	case "all":
		fmt.Fprintf(ser.w, ser.us.GetAllUserData())
	case "login":
		val, exi := args["account"]
		val2, exi2 := args["password"]

		if exi && exi2 {
			fmt.Fprint(ser.w, ser.us.Login(val[0], val2[0]))
		} else {
			fmt.Fprint(ser.w, "argument error")
		}
	case "delete":
		val, exi := args["account"]
		val2, exi2 := args["password"]

		if exi && exi2 {
			fmt.Fprint(ser.w, ser.us.DeleteUser(val[0], val2[0]))
		} else {
			fmt.Fprint(ser.w, "argument error")
		}
	case "regist":
		val, exi := args["account"]
		val2, exi2 := args["password"]
		val3, exi3 := args["name"]

		if exi && exi2 && exi3 {
			fmt.Fprint(ser.w, ser.us.Regist(val[0], val2[0], val3[0]))
		} else {
			fmt.Fprint(ser.w, "argument error")
		}
	default:
		http.NotFound(ser.w, ser.r)
	}
}

func (ser *server) fetchProduct(path []string, args map[string][]string) {
	if len(path) != 2 {
		http.NotFound(ser.w, ser.r)
		return
	}

	switch path[1] {
	case "all":
		fmt.Fprint(ser.w, ser.pr.GetAll())
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
				fmt.Fprint(ser.w, ser.pr.AddProduct(name[0], p, des[0], a, account[0], b, date[0]))
			} else {
				fmt.Fprint(ser.w, "price or amount was not an integer.")
			}
		} else {
			fmt.Fprint(ser.w, "argument error")
		}
	case "delete":
	case "newproduct":
		val, exi := args["amount"]

		if exi {
			v, err := strconv.Atoi(val[0])

			if err == nil {
				fmt.Fprint(ser.w, ser.pr.GetNewest(v))
			} else {
				fmt.Fprint(ser.w, "amount was not an integer.")
			}

		} else {
			fmt.Fprint(ser.w, "argument error")
		}
	case "search":
		val, exi := args["name"]

		if exi {
			fmt.Fprint(ser.w, ser.pr.SearchProductsByName(val[0]))
		} else {
			fmt.Fprint(ser.w, "argument error")
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
				fmt.Fprint(ser.w, ser.pr.EnhanceSearchProductsByName(name[0], mi, ma, ev))
			} else {
				fmt.Fprint(ser.w, "min price, max price or evaluation was not as interger")
			}
		} else {
			fmt.Fprint(ser.w, "argument error")
		}
	default:
		http.NotFound(ser.w, ser.r)
	}
}

func (ser *server) help() {
	fmt.Fprintln(ser.w, `
		<html>
			<p> 
				/user/all<br>
				列出所有帳號(僅限開發期間)<br>
				<a href=/user/all> /user/all </a><br><br>
			</p>
			<p> 
				/history/add?account=...&password=..pdidb1			登入是否成功(bool)<br>
				e.g.登入帳號為test@gmail.com以及密碼為0000的使用者<br>
				<a href=https://se-ssb.herokuapp.com/history/add?account=test@gmail.com&pdid=1>
				https://se-ssb.herokuapp.com/history/add?account=test@gmail.com&pdid=1</a>
				<br><br>
			</p>
			<p>
				/user/regist?account=...&password=...&name=...<br>
				註冊新帳號<br>
				e.g.註冊一帳號為test2@gmail.com，密碼為1234，使用者姓名為Wilson的帳號<br>
				<a href=https://se-ssb.herokuapp.com/user/regist?account=test2@gmail.com&password=1234&name=Wilson>
				https://se-ssb.herokuapp.com/user/regist?account=test2@gmail.com&password=1234&name=Wilson</a>
				<br><br>
			</p>
			<p>
				/user/delete?account=...&password=...<br>
				刪除帳號<br>
				e.g.刪除帳號為test3@gmail.com的帳號(需要輸入密碼驗證:密碼為1234)<br>
				<a href=https://se-ssb.herokuapp.com/user/delete?account=test3@gmail.com&password=1234>
				https://se-ssb.herokuapp.com/user/delete?account=test3@gmail.com&password=1234</a>
				<br><br>
			</p>
			<p> 
				/product/all<br>
				列出所有商品(僅限開發期間)<br>
				<a href=/product/all> /product/all </a><br><br>
			</p>
			<p>
				/product/newest?amount=...<br>
				e.g.顯示最新商品(3筆資料)<br>
				<a href=https://se-ssb.herokuapp.com/product?amount=3>
				https://se-ssb.herokuapp.com/product?amount=3</a>
				<br><br>
			<p> 
				/product/add?name=...&price=...&description=...&amount=...&account=...&bid=...&date=...<br>
				新增商品<br>
				e.g.新增一商品->商品名:ifone12價格:5000，商品說明:盜版商品，帳號:test@gmail.com，競標:是，競標期限:2020-12-31<br>
				<a href=https://se-ssb.herokuapp.com/product/add?name=ifone12&price=5000&description=盜版商品&amount=10&account=test@gmail.com&bid=true&date=2020-12-31>
				https://se-ssb.herokuapp.com/product/add?name=ifone12&price=5000&description=盜版商品&amount=10&account=test@gmail.com&bid=true&date=2020-12-31</a>
				<br><br>
			</p>
			<p>
				/product/search?name=...<br>
				查詢商品<br>
				e.g.查詢商品名中含有"ifone"的商品<br>
				<a href=https://se-ssb.herokuapp.com/product/search?name=ifone>
				https://se-ssb.herokuapp.com/product/search?name=ifone</a>
				<br><br>
			</p>
			<p>
				/product/search?name=...&minprice=...&maxprice=...&eval=...<br>
				查詢商品(過濾)<br>
				e.g.查詢商品名中含有"ifone"的商品，最低價格為10，最高價格為5000，最低評價為2<br>
				<a href=https://se-ssb.herokuapp.com/product/search?name=ifone&minprice=10&maxprice=5000&eval=2>
				https://se-ssb.herokuapp.com/product/search?name=ifone&minprice=10&maxprice=5000&eval=2</a>
				<br><br>
			</p>
			<p> 
				/history/all<br>
				列出歷史紀錄(僅限開發期間)<br>
				<a href=/user/all> /user/all </a><br><br>
			</p>
			<p> 
				/history/add?account=...&pdid=...<br>
				增加一筆新的歷史紀錄<br>
				e.g.新增帳號為test@gmail.com以及商品id為1的歷史紀錄<br>
				<a href=https://se-ssb.herokuapp.com/history/add?account=test@gmail.com&pdid=1>
				https://se-ssb.herokuapp.com/history/add?account=test@gmail.com&pdid=1</a>
				<br><br>
			</p>
			<p>
				/history/search?account=...&amount=...<br>
				查詢歷史紀錄<br>
				e.g.查詢帳號為test2@gmail.com的10歷史紀錄<br>
				<a href=https://se-ssb.herokuapp.com/history/search?amount=10>
				https://se-ssb.herokuapp.com/history/search?amount=10</a>
				<br><br>
			</p>
			<p>
				/history/delete?account=...&pdid=...<br>
				刪除歷史紀錄<br>
				e.g.刪除帳號test3@gmail.com以及商品編號為2的歷史紀錄<br>
				<a href=https://se-ssb.herokuapp.com/history/delete?account=test3@gmail.com&pdid=2>
				https://se-ssb.herokuapp.com/history/delete?account=test3@gmail.com&pdid=2</a>
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
