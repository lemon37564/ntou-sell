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

func (ser Server) defaultFunc(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	fmt.Fprintln(w, API)
}

func (ser Server) fetchBid(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, BidAPI)

	case "get": //For single bid product
		pdid := args.Get("pdid")
		res, err := ser.Bd.GetProductBidInfo(pdid)

		if err == nil {
			fmt.Fprint(w, res)
		} else {
			http.Error(w, res, http.StatusBadRequest)
		}

	case "set":
		pdid := args.Get("pdid")
		money := args.Get("money")

		res, err := ser.Bd.SetBidForBuyer(uid, pdid, money)
		fmt.Fprint(w, res)
		if err != nil {
			log.Println(err)
		}

	case "delete":
		pdid := args.Get("pdid")

		res, err := ser.Bd.DeleteBid(pdid)
		fmt.Fprint(w, res)
		if err != nil {
			log.Println(err)
		}

	default:
		http.NotFound(w, r)
	}
}

func (ser Server) fetchCart(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, CartAPI)

	case "add": //For single product
		pdid, err := strconv.Atoi(args.Get("pdid"))
		amount, err2 := strconv.Atoi(args.Get("amount"))

		if err == nil && err2 == nil {
			fmt.Fprint(w, ser.Ct.AddProductToCart(uid, pdid, amount))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	case "remo":
		pdid, err := strconv.Atoi(args.Get("pdid"))

		if err == nil {
			fmt.Fprint(w, ser.Ct.RemoveProduct(uid, pdid))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	case "modf":
		pdid, err := strconv.Atoi(args.Get("pdid"))
		amount, err2 := strconv.Atoi(args.Get("amount"))

		if err == nil && err2 == nil {
			fmt.Fprint(w, ser.Ct.ModifyAmount(uid, pdid, amount))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	// case "tal":
	// 	fmt.Fprint(w, ser.Ct.TotalCount(uid))

	case "geps": //拿商品
		fmt.Fprint(w, ser.Ct.GetProducts(uid))

	default:
		http.NotFound(w, r)
	}
}

func (ser Server) fetchHistory(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, HistoryAPI)

	case "get":
		amount, err := strconv.Atoi(args.Get("amount"))
		newest := (args.Get("newest") == "true")

		if err == nil {
			fmt.Fprint(w, ser.Ht.GetHistory(uid, amount, newest))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	case "delete":
		pdid, err := strconv.Atoi(args.Get("pdid"))

		if err == nil {
			fmt.Fprint(w, ser.Ht.Delete(uid, pdid))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	case "deleteall":
		fmt.Fprint(w, ser.Ht.DeleteAll(uid))

	case "deletespec":
		pdids := args.Get("pdids")
		fmt.Fprint(w, ser.Ht.DeleteSpecific(uid, pdids))

	case "add":
		pdid, err := strconv.Atoi(args.Get("pdid"))

		if err == nil {
			fmt.Fprint(w, ser.Ht.AddHistory(uid, pdid))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	default:
		http.NotFound(w, r)
	}
}

func (ser Server) fetchMessage(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, MessageAPI)

	case "all":
		fmt.Fprint(w, ser.Ms.GetAll())

	case "send":
		ruid, err := strconv.Atoi(args.Get("remoteUID"))
		val2 := args.Get("text")

		if err == nil {
			fmt.Fprint(w, ser.Ms.AddMessage(uid, ruid, val2))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	case "get":
		ruid, err := strconv.Atoi(args.Get("remoteUID"))
		asc := (args.Get("ascend") == "true")

		if err == nil {
			fmt.Fprint(w, ser.Ms.GetMessages(uid, ruid, asc))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	default:
		http.NotFound(w, r)
	}
}

func (ser Server) fetchOrder(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, OrderAPI)

	case "get":
		fmt.Fprint(w, ser.Od.GetOrders(uid))

	case "add":
		pdid, err := strconv.Atoi(args.Get("pdid"))
		amount, err2 := strconv.Atoi(args.Get("amount"))

		if err == nil && err2 == nil {
			fmt.Fprint(w, ser.Od.AddOrder(uid, pdid, amount))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	case "del":
		pdid, err := strconv.Atoi(args.Get("pdid"))
		if err == nil {
			fmt.Fprint(w, ser.Od.Delete(uid, pdid))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	default:
		http.NotFound(w, r)
	}

}

func (ser Server) fetchProduct(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	uid, valid := sessionValid(w, r)
	if !valid {
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

			fmt.Fprint(w, "ok")

			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				log.Println("at upload file:", err)
				return
			}
			defer file.Close()

			spt := strings.Split(handler.Filename, ".")
			subName := spt[len(spt)-1]

			fmt.Fprint(w, handler.Header)
			f, err := os.Create("webpage/img/" + fmt.Sprint(pdid) + subName)
			if err != nil {
				log.Println("at creatingfile:", err)
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
		fmt.Fprint(w, ProductAPI)

	case "add":
		name = args.Get("name")
		price = args.Get("price")
		des = args.Get("description")
		amount = args.Get("amount")
		bid = args.Get("bid")
		date = args.Get("date")

		p, err := strconv.Atoi(price)
		a, err2 := strconv.Atoi(amount)
		b := (bid == "true")

		if err == nil && err2 == nil {
			_, stat := ser.Pd.AddProduct(name, p, des, a, uid, b, date)
			fmt.Fprint(w, stat)
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	case "get":
		pdid, err := strconv.Atoi(args.Get("pdid"))

		if err == nil {
			fmt.Fprint(w, ser.Pd.GetProductInfo(pdid))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	case "delete":
		val := args.Get("pdname")

		if exi {
			fmt.Fprint(w, ser.Pd.DeleteProduct(uid, val))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	case "newest":
		amount, err := strconv.Atoi(args.Get("amount"))

		if err == nil {
			fmt.Fprint(w, ser.Pd.GetNewest(amount))
		} else {
			http.Error(w, "argument error", http.StatusBadRequest)
		}

	case "search":
		name := args.Get("name")
		min := args.Get("min")
		max := args.Get("max")
		eval := args.Get("eval")

		// fmt.Fprint(w, ser.Pd.EnhanceSearchProductsByName(name, min, max, eval))

		// if exist0 && (err1 != nil && err2 != nil && err3 != nil) {
		// 	fmt.Fprint(w, ser.Pd.SearchProductsByName(name))
		// } else if exist0 && (err1 == nil || err2 == nil || err3 == nil) {
		// 	fmt.Fprint(w, ser.Pd.EnhanceSearchProductsByName(name, min, max, eval))
		// } else {
		// 	http.Error(w, "argument error", http.StatusBadRequest)
		// }

	case "urproduct":
		fmt.Fprint(w, ser.Pd.GetSellerProduct(uid))

	default:
		http.NotFound(w, r)
	}
}

func (ser Server) fetchUser(w http.ResponseWriter, r *http.Request) {
	if !ser.validation(w, r) {
		return
	}

	// POST
	path := mux.Vars(r)
	r.ParseForm()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, UserAPI)

	case "login":
		uid, valid := ser.Ur.Login(r.Form["account"][0], r.Form["password"][0])

		if valid {
			// set session to maintain login condition
			login(w, r, uid)
			fmt.Fprintln(w, "登入成功!")
		} else {
			fmt.Fprint(w, "登入失敗")
		}

	case "regist":
		fmt.Fprint(w, ser.Ur.Regist(r.Form["account"][0], r.Form["password"][0], r.Form["name"][0]))

	default:
		uid, valid := sessionValid(w, r)
		if valid {
			switch path["key"] {
			case "delete":
				fmt.Fprint(w, ser.Ur.DeleteUser(uid, r.Form["password"][0]))

			case "changePassword":
				fmt.Fprint(w, ser.Ur.ChangePassword(uid, r.Form["oldPassword"][0], r.Form["newPassword"][0]))

			case "changeName":
				fmt.Fprint(w, ser.Ur.ChangeName(uid, r.Form["newName"][0]))

			case "logout":
				logout(w, r)
				fmt.Fprintln(w, "已登出")

			default:
				http.NotFound(w, r)
			}
		}
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
