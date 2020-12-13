package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ser *Server) defaultFunc(w http.ResponseWriter, r *http.Request) {
	switch mux.Vars(r)["key"] {
	case "success":
		fmt.Fprintln(w, "登入成功!")
	case "testpic":
		fmt.Fprint(w, `<html><img src="https://se-ssb.herokuapp.com/backend/pics/server.jpg"></html>`)
	default:
		fmt.Fprintln(w, helpPage)
	}
}

func (ser *Server) fetchHistory(w http.ResponseWriter, r *http.Request) {
	if !sessionValid(w, r) {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
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
		pdid, exi2 := args["pdid"]

		if exi && exi2 {
			pd, err := strconv.Atoi(pdid[0])
			if err == nil {
				uid := ser.Ur.GetUIDFromAccount(ac[0])
				fmt.Fprint(w, ser.Ht.AddHistory(uid, pd))
			} else {
				fmt.Fprint(w, "pdid was not an integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) fetchUser(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "login":

		account, exi := args["account"]
		pass, exi2 := args["password"]

		if exi && exi2 {
			valid := ser.Ur.Login(account[0], pass[0])

			// set cookies to maintain login condition
			if valid {
				login(w, r)
				fmt.Fprintln(w, "登入成功!")
			} else {
				fmt.Fprint(w, "登入失敗")
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
	case "logout":
		logout(w, r)
		fmt.Fprintln(w, "已登出")
	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) fetchProduct(w http.ResponseWriter, r *http.Request) {
	if !sessionValid(w, r) {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
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
				_, stat := ser.Pd.AddProduct(name[0], p, des[0], a, account[0], b, date[0])
				fmt.Fprint(w, stat)
			} else {
				fmt.Fprint(w, "price or amount was not an integer.")
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
	case "filterSearch":
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
				fmt.Fprint(w, "min price, max price or evaluation was not as interger")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) fetchOrder(w http.ResponseWriter, r *http.Request) {
	if !sessionValid(w, r) {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "get":
		uid, ex1 := args["uid"]

		if ex1 {
			i, err1 := strconv.Atoi(uid[0])
			if err1 == nil {
				fmt.Fprint(w, ser.Od.GetOrders(i))
			} else {
				fmt.Fprint(w, "User id not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "add":
		exist := make([]bool, 3)
		var uid, pdid, amount []string

		uid, exist[0] = args["uid"]
		pdid, exist[1] = args["pdid"]
		amount, exist[2] = args["amount"]
		if all(exist) {
			ui, err1 := strconv.Atoi(uid[0])
			pi, err2 := strconv.Atoi(pdid[0])
			amo, err3 := strconv.Atoi(amount[0])

			if err1 == nil && err2 == nil && err3 == nil {
				fmt.Fprint(w, ser.Od.AddOrder(ui, pi, amo))
			} else {
				fmt.Fprint(w, "Userid,Productid or amount was not as interger")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "del":
		uid, exi := args["uid"]
		pdid, exi2 := args["pdid"]
		if exi && exi2 {
			ui, err := strconv.Atoi(uid[0])
			pi, err1 := strconv.Atoi(pdid[0])
			if err == nil && err1 == nil {
				fmt.Fprint(w, ser.Od.Delete(ui, pi))
			} else {
				fmt.Fprint(w, "User id or Product id was not an integer.")
			}

		} else {
			fmt.Fprint(w, "argument error")
		}
	default:
		http.NotFound(w, r)
	}

}

func (ser *Server) fetchBid(w http.ResponseWriter, r *http.Request) {
	if !sessionValid(w, r) {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "get": //For single bid product
		pdid, ex1 := args["pdid"]

		if ex1 {
			i, err1 := strconv.Atoi(pdid[0])
			if err1 == nil {

				fmt.Fprint(w, ser.Bd.GetProductBidInfo(i))
			} else {
				fmt.Fprint(w, "product id not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "set":
		exist := make([]bool, 3)
		var pdid, uid, money []string

		pdid, exist[0] = args["pdid"]
		uid, exist[1] = args["uid"]
		money, exist[2] = args["money"]

		if all(exist) {
			p, err1 := strconv.Atoi(pdid[0])
			u, err2 := strconv.Atoi(uid[0])
			m, err3 := strconv.Atoi(money[0])

			if err1 == nil && err2 == nil && err3 == nil {
				fmt.Fprint(w, ser.Bd.SetBidForBuyer(p, u, m))
			} else {
				fmt.Fprint(w, "User, Product id or money was not an integer.")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "delete":
		pdid, ex1 := args["pdid"]

		if ex1 {
			p, err1 := strconv.Atoi(pdid[0])
			if err1 == nil {
				fmt.Fprint(w, ser.Bd.DeleteBid(p))
			} else {
				fmt.Fprint(w, "Product id not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}

	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) fetchCart(w http.ResponseWriter, r *http.Request) {
	if !sessionValid(w, r) {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "add": //For single product
		uid, ex1 := args["uid"]
		pdid, ex2 := args["pdid"]
		amount, ex3 := args["amount"]

		if ex1 && ex2 && ex3 {
			u, err1 := strconv.Atoi(uid[0])
			p, err2 := strconv.Atoi(pdid[0])
			amo, err3 := strconv.Atoi(amount[0])
			if err1 == nil && err2 == nil && err3 == nil {
				fmt.Fprint(w, ser.Ct.AddProductToCart(u, p, amo))

			} else {
				fmt.Fprint(w, "User id ,product id or amount was not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "remo":
		uid, ex1 := args["uid"]
		pdid, ex2 := args["pdid"]

		if ex1 && ex2 {
			u, err1 := strconv.Atoi(uid[0])
			p, err2 := strconv.Atoi(pdid[0])

			if err1 == nil && err2 == nil {
				fmt.Fprint(w, ser.Ct.RemoveProduct(u, p))

			} else {
				fmt.Fprint(w, "User id ,product id was not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "modf":
		uid, ex1 := args["uid"]
		pdid, ex2 := args["pdid"]
		amount, ex3 := args["amount"]

		if ex1 && ex2 && ex3 {
			u, err1 := strconv.Atoi(uid[0])
			p, err2 := strconv.Atoi(pdid[0])
			amo, err3 := strconv.Atoi(amount[0])
			if err1 == nil && err2 == nil && err3 == nil {
				fmt.Fprint(w, ser.Ct.ModifyAmount(u, p, amo))

			} else {
				fmt.Fprint(w, "User id ,product id or amount was not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "tal":
		uid, ex1 := args["uid"]

		if ex1 {
			u, err1 := strconv.Atoi(uid[0])

			if err1 == nil {
				fmt.Fprint(w, ser.Ct.TotalCount(u))

			} else {
				fmt.Fprint(w, "User id was not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "geps": //拿商品
		uid, ex1 := args["uid"]

		if ex1 {
			u, err1 := strconv.Atoi(uid[0])
			if err1 == nil {
				fmt.Fprint(w, ser.Ct.GetProducts(u))

			} else {
				fmt.Fprint(w, "User id was not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) fetchSell(w http.ResponseWriter, r *http.Request) {
	if !sessionValid(w, r) {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "set": //For single bid product
		exist := make([]bool, 9)
		var pdname, price, description, amount, account, sellerID, bid, date, dateLine []string
		pdname, exist[0] = args["pdname"]
		price, exist[1] = args["price"]
		description, exist[2] = args["description"]
		amount, exist[3] = args["amount"]
		account, exist[4] = args["account"]
		sellerID, exist[5] = args["sellerID"]
		bid, exist[6] = args["bid"]
		date, exist[7] = args["date"]
		dateLine, exist[8] = args["dateLine"]

		if all(exist) {
			pr, err1 := strconv.Atoi(price[0])
			amo, err2 := strconv.Atoi(amount[0])
			sel, err3 := strconv.Atoi(sellerID[0])
			var bi bool
			if bid[0] == "true" {
				bi = true
			}

			if err1 == nil && err2 == nil && err3 == nil {
				fmt.Fprint(w, ser.Se.SetProductpdid(pdname[0], pr, description[0], amo, account[0], sel, bi, date[0], dateLine[0]))
			} else {
				fmt.Fprint(w, "data has wrong")
			}

		} else {
			fmt.Fprint(w, "argument error")
		}

	default:
		http.NotFound(w, r)
	}
}

func all(bs []bool) bool {
	for _, v := range bs {
		if !v {
			return false
		}
	}

	return true
}
