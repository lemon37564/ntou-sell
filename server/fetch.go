package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (ser *Server) fetch(w http.ResponseWriter, r *http.Request, cmd string, args map[string][]string) {
	path := strings.Split(cmd, "/")

	if len(path) == 0 {
		http.NotFound(w, r)
	} else if len(path) == 1 && path[0] == "help" {
		ser.help(w, r)
	} else if path[0] == "user" {
		// user functions need to be in front of verification, or no one can log in anymore
		ser.fetchUser(w, r, path, args)
	} else {
		if !ser.verify(w, r) {
			fmt.Fprint(w, "請先登入!!")
			return
		}

		switch path[0] {
		case "product":
			ser.fetchProduct(w, r, path, args)
		case "history":
			ser.fetchHistory(w, r, path, args)
		default:
			http.NotFound(w, r)
		}
	}
}

func (ser *Server) fetchHistory(w http.ResponseWriter, r *http.Request, path []string, args map[string][]string) {
	if len(path) != 2 {
		http.NotFound(w, r)
		return
	}

	switch path[1] {
	case "all":
		fmt.Fprint(w, ser.Ht.GetAll())
	case "get":
		ac, exi := args["account"]
		val, exist := args["amount"]

		if exist && exi {
			amnt, err := strconv.Atoi(val[0])
			if err == nil {
				uid := ser.Ur.GetUIDFromAccount(ac[0])
				fmt.Fprint(w, ser.Ht.GetHistory(uid, amnt))
			} else {
				fmt.Fprint(w, "amount was not an integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "delete":
		ac, exi := args["account"]
		pdid, exi2 := args["pdid"]

		if exi && exi2 {
			pd, err := strconv.Atoi(pdid[0])
			if err == nil {
				uid := ser.Ur.GetUIDFromAccount(ac[0])
				fmt.Fprint(w, ser.Ht.GetHistory(uid, pd))
			} else {
				fmt.Fprint(w, "pdid was not an integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "add":
		ac, exi := args["account"]
		pdid, exi2 := args["pd_id"]

		if exi && exi2 {
			pd, err := strconv.Atoi(pdid[0])
			if err == nil {
				uid := ser.Ur.GetUIDFromAccount(ac[0])
				fmt.Fprint(w, ser.Ht.AddHistory(uid, pd))
			} else {
				fmt.Fprint(w, "pd_id was not an integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) fetchUser(w http.ResponseWriter, r *http.Request, path []string, args map[string][]string) {
	if len(path) != 2 {
		http.NotFound(w, r)
		return
	}

	switch path[1] {
	case "all":
		fmt.Fprintf(w, ser.Ur.GetAllUserData())
	case "login":
		account, exi := args["account"]
		pass, exi2 := args["password"]

		if exi && exi2 {
			valid := ser.Ur.Login(account[0], pass[0])
			fmt.Fprint(w, valid)

			// set cookies to maintain login condition
			if valid {
				ser.setCookies(w, r, account[0], pass[0])
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "delete":
		val, exi := args["account"]
		val2, exi2 := args["password"]

		if exi && exi2 {
			fmt.Fprint(w, ser.Ur.DeleteUser(val[0], val2[0]))
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "regist":
		val, exi := args["account"]
		val2, exi2 := args["password"]
		val3, exi3 := args["name"]

		if exi && exi2 && exi3 {
			fmt.Fprint(w, ser.Ur.Regist(val[0], val2[0], val3[0]))
		} else {
			fmt.Fprint(w, "argument error")
		}
	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) fetchProduct(w http.ResponseWriter, r *http.Request, path []string, args map[string][]string) {
	if len(path) != 2 {
		http.NotFound(w, r)
		return
	}

	switch path[1] {
	case "all":
		fmt.Fprint(w, ser.Pd.GetAll())
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
				fmt.Fprint(w, ser.Pd.AddProduct(name[0], p, des[0], a, account[0], b, date[0]))
			} else {
				fmt.Fprint(w, "price Od amount was not an integer.")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "delete":
	case "newest":
		val, exi := args["amount"]

		if exi {
			v, err := strconv.Atoi(val[0])

			if err == nil {
				fmt.Fprint(w, ser.Pd.GetNewest(v))
			} else {
				fmt.Fprint(w, "amount was not an integer.")
			}

		} else {
			fmt.Fprint(w, "argument error")
		}
	case "search":
		val, exi := args["name"]

		if exi {
			fmt.Fprint(w, ser.Pd.SearchProductsByName(val[0]))
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
				fmt.Fprint(w, ser.Pd.EnhanceSearchProductsByName(name[0], mi, ma, ev))
			} else {
				fmt.Fprint(w, "min price, max price Od evaluation was not as interger")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
		<html>
			<H1>後端API</H1>
			<H6>測試用帳密:account=1234&password=1234</H6>
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
				<a href=https://se-ssb.herokuapp.com/product/newest?amount=3>
				https://se-ssb.herokuapp.com/product/newest?amount=3</a>
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
				<a href=/history/all> /history/all </a><br><br>
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
				/history/get?account=...&amount=...<br>
				查詢歷史紀錄<br>
				e.g.查詢帳號為test2@gmail.com的10歷史紀錄<br>
				<a href=https://se-ssb.herokuapp.com/history/get?account=test2@gmail.com&amount=10>
				https://se-ssb.herokuapp.com/history/get?account=test2@gmail.com&amount=10</a>
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
