package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (ser *Server) defaultFunc(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	fmt.Fprintln(w, HelpPage)
}

func (ser *Server) fetchHistory(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, HistoryHelp)
	case "get":
		val, exist := args["amount"]
		val2, exi2 := args["newest"]

		if exist && exi2 {
			amnt, err := strconv.Atoi(val[0])
			if err == nil {
				fmt.Fprint(w, ser.Ht.GetHistory(uid, amnt, val2[0] == "true"))
			} else {
				fmt.Fprint(w, "amount was not an integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "delete":
		pdid, exi := args["pdid"]

		if exi {
			pd, err := strconv.Atoi(pdid[0])
			if err == nil {
				fmt.Fprint(w, ser.Ht.Delete(uid, pd))
			} else {
				fmt.Fprint(w, "pdid was not an integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "add":
		pdid, exi := args["pdid"]

		if exi {
			pd, err := strconv.Atoi(pdid[0])
			if err == nil {
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
	if !ser.validation(w, r) {
		return
	}

	path := mux.Vars(r)
	r.ParseForm()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, UserHelp)
	case "login":

		uid, valid := ser.Ur.Login(r.Form["account"][0], r.Form["password"][0])

		if valid {
			// set session to maintain login condition
			login(w, r, uid)
			fmt.Fprintln(w, "登入成功!")
		} else {
			fmt.Fprint(w, "登入失敗")
		}

	case "delete":
		fmt.Fprint(w, ser.Ur.DeleteUser(r.Form["account"][0], r.Form["password"][0]))

	case "regist":
		fmt.Fprint(w, ser.Ur.Regist(r.Form["account"][0], r.Form["password"][0], r.Form["name"][0]))

	case "changePassword":
		fmt.Fprint(w, ser.Ur.ChangePassword(r.Form["account"][0], r.Form["oldPassword"][0], r.Form["newPassword"][0]))

	case "changeName":
		fmt.Fprint(w, ser.Ur.ChangeName(r.Form["account"][0], r.Form["newName"][0]))

	case "logout":
		logout(w, r)
		fmt.Fprintln(w, "已登出")
	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) fetchProduct(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		fmt.Fprint(w, "請先登入!")
		return
	}

	// temp area
	if r.Method == "POST" {
		log.Println("receive post (product)")
		if mux.Vars(r)["key"] == "postadd" {
			var pdid int
			var stat string

			r.ParseMultipartForm(32 << 20)

			name := r.Form["name"][0]
			price := r.Form["price"][0]
			des := r.Form["description"][0]
			amount := r.Form["amount"][0]
			bid := r.Form["bid"][0]
			date := r.Form["date"][0]

			p, err1 := strconv.Atoi(price)
			a, err2 := strconv.Atoi(amount)
			b := (bid == "true")

			if err1 == nil && err2 == nil {
				pdid, stat = ser.Pd.AddProduct(name, p, des, a, uid, b, date)
			}

			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				log.Println(err)
				return
			}
			defer file.Close()

			spt := strings.Split(handler.Filename, ".")
			subName := spt[len(spt)-1]

			fmt.Fprint(w, handler.Header)
			f, err := os.Create("webpage/img/" + fmt.Sprint(pdid) + subName)
			if err != nil {
				log.Println(err)
				return
			}
			defer f.Close()

			io.Copy(f, file)
			fmt.Fprint(w, "\n", stat)
		}
		return
	}
	// temp area

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, ProductHelp)
	case "all":
		fmt.Fprint(w, ser.Pd.GetAll())
	case "add":
		exist := make([]bool, 6)
		var name, price, des, amount, bid, date []string

		name, exist[0] = args["name"]
		price, exist[1] = args["price"]
		des, exist[2] = args["description"]
		amount, exist[3] = args["amount"]
		bid, exist[4] = args["bid"]
		date, exist[5] = args["date"]

		if all(exist) {
			p, err1 := strconv.Atoi(price[0])
			a, err2 := strconv.Atoi(amount[0])
			b := (bid[0] == "true")

			if err1 == nil && err2 == nil {
				_, stat := ser.Pd.AddProduct(name[0], p, des[0], a, uid, b, date[0])
				fmt.Fprint(w, stat)
			} else {
				fmt.Fprint(w, "price or amount was not an integer.")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "get":
		val, exi := args["pdid"]

		if exi {
			if pdid, err := strconv.Atoi(val[0]); err == nil {
				fmt.Fprint(w, ser.Pd.GetProductInfo(pdid))
			} else {
				fmt.Fprint(w, "pdid is not an integer!")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "delete":
		val, exi := args["pdname"]

		if exi {
			fmt.Fprint(w, ser.Pd.DeleteProduct(uid, val[0]))
		} else {
			fmt.Fprint(w, "argument error")
		}
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
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, OrderHelp)
	case "get":
		fmt.Fprint(w, ser.Od.GetOrders(uid))

	case "add":
		pdid, exi := args["pdid"]
		amount, exi2 := args["amount"]
		if exi && exi2 {
			pi, err2 := strconv.Atoi(pdid[0])
			amo, err3 := strconv.Atoi(amount[0])

			if err2 == nil && err3 == nil {
				fmt.Fprint(w, ser.Od.AddOrder(uid, pi, amo))
			} else {
				fmt.Fprint(w, "Userid,Productid or amount was not as interger")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "del":
		pdid, exi := args["pdid"]
		if exi {
			if pi, err1 := strconv.Atoi(pdid[0]); err1 == nil {
				fmt.Fprint(w, ser.Od.Delete(uid, pi))
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
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, BidHelp)
	case "get": //For single bid product
		pdid, ex1 := args["pdid"]

		if ex1 {
			if i, err1 := strconv.Atoi(pdid[0]); err1 == nil {
				fmt.Fprint(w, ser.Bd.GetProductBidInfo(i))
			} else {
				fmt.Fprint(w, "product id not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "set":
		pdid, exi := args["pdid"]
		money, exi2 := args["money"]

		if exi && exi2 {
			p, err1 := strconv.Atoi(pdid[0])
			m, err3 := strconv.Atoi(money[0])

			if err1 == nil && err3 == nil {
				fmt.Fprint(w, ser.Bd.SetBidForBuyer(p, uid, m))
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
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, CartHelp)
	case "all":
		fmt.Fprint(w, ser.Ct.Debug())
	case "add": //For single product
		pdid, ex2 := args["pdid"]
		amount, ex3 := args["amount"]

		if ex2 && ex3 {
			p, err2 := strconv.Atoi(pdid[0])
			amo, err3 := strconv.Atoi(amount[0])
			if err2 == nil && err3 == nil {
				fmt.Fprint(w, ser.Ct.AddProductToCart(uid, p, amo))

			} else {
				fmt.Fprint(w, "product id or amount was not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "remo":
		pdid, exi := args["pdid"]

		if exi {
			if p, err := strconv.Atoi(pdid[0]); err == nil {
				fmt.Fprint(w, ser.Ct.RemoveProduct(uid, p))

			} else {
				fmt.Fprint(w, "product id was not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "modf":
		pdid, ex2 := args["pdid"]
		amount, ex3 := args["amount"]

		if ex2 && ex3 {
			p, err2 := strconv.Atoi(pdid[0])
			amo, err3 := strconv.Atoi(amount[0])
			if err2 == nil && err3 == nil {
				fmt.Fprint(w, ser.Ct.ModifyAmount(uid, p, amo))

			} else {
				fmt.Fprint(w, "User id ,product id or amount was not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	// case "tal":
	// 	fmt.Fprint(w, ser.Ct.TotalCount(uid))

	case "geps": //拿商品
		fmt.Fprint(w, ser.Ct.GetProducts(uid))

	default:
		http.NotFound(w, r)
	}
}

func (ser *Server) fetchMessage(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		fmt.Fprint(w, "請先登入!")
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, MessageHelp)
	case "all":
		fmt.Fprint(w, ser.Ms.GetAll())
	case "send":
		val, exi := args["remoteUID"]
		val2, exi2 := args["text"]

		if exi && exi2 {
			ruid, err := strconv.Atoi(val[0])
			if err == nil {
				fmt.Fprint(w, ser.Ms.AddMessage(uid, ruid, val2[0]))
			} else {
				fmt.Fprint(w, "receiverUID is not integer")
			}
		} else {
			fmt.Fprint(w, "argument error")
		}
	case "get":
		val, exi := args["remoteUID"]
		val2, exi2 := args["ascend"]

		if exi && exi2 {
			ruid, err := strconv.Atoi(val[0])
			if err == nil {
				fmt.Fprint(w, ser.Ms.GetMessages(uid, ruid, val2[0] == "true"))
			} else {
				fmt.Fprint(w, "receiverUID is not integer")
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
