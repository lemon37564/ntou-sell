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
		response(w, res, err)

	case "set":
		pdid := args.Get("pdid")
		money := args.Get("money")

		res, err := ser.Bd.SetBidForBuyer(uid, pdid, money)
		response(w, res, err)

	case "delete":
		pdid := args.Get("pdid")

		res, err := ser.Bd.DeleteBid(pdid)
		response(w, res, err)

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
		pdid := args.Get("pdid")
		amount := args.Get("amount")

		res, err := ser.Ct.AddProductToCart(uid, pdid, amount)
		response(w, res, err)

	case "remo":
		pdid := args.Get("pdid")

		res, err := ser.Ct.RemoveProduct(uid, pdid)
		response(w, res, err)

	case "modf":
		pdid := args.Get("pdid")
		amount := args.Get("amount")

		res, err := ser.Ct.ModifyAmount(uid, pdid, amount)
		response(w, res, err)

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
		amount := args.Get("amount")
		newest := args.Get("newest")

		res, err := ser.Ht.GetHistory(uid, amount, newest)
		response(w, res, err)

	case "delete":
		pdid := args.Get("pdid")

		res, err := ser.Ht.Delete(uid, pdid)
		response(w, res, err)

	case "deleteall":
		fmt.Fprint(w, ser.Ht.DeleteAll(uid))

	case "deletespec":
		pdids := args.Get("pdids")
		res, err := ser.Ht.DeleteSpecific(uid, pdids)
		response(w, res, err)

	case "add":
		pdid := args.Get("pdid")

		res, err := ser.Ht.AddHistory(uid, pdid)
		response(w, res, err)

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
		ruid := args.Get("remoteUID")
		txt := args.Get("text")

		res, err := ser.Ms.AddMessage(uid, ruid, txt)
		response(w, res, err)

	case "get":
		ruid := args.Get("remoteUID")
		asc := args.Get("ascend")

		res, err := ser.Ms.GetMessages(uid, ruid, asc)
		response(w, res, err)

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
		pdid := args.Get("pdid")
		amount := args.Get("amount")

		res, err := ser.Od.AddOrder(uid, pdid, amount)
		response(w, res, err)

	case "del":
		pdid := args.Get("pdid")

		res, err := ser.Od.Delete(uid, pdid)
		response(w, res, err)

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

	if r.Method == "POST" {
		if mux.Vars(r)["key"] == "postadd" {
			r.ParseMultipartForm(32 << 20)

			name := r.FormValue("name")
			price := r.FormValue("price")
			des := r.FormValue("description")
			amount := r.FormValue("amount")
			bid := r.FormValue("bid")
			date := r.FormValue("date")

			pdid, err := ser.Pd.AddProduct(uid, name, price, des, amount, bid, date)
			if err == nil {
				fmt.Fprint(w, "ok")
			} else {
				http.Error(w, "failed", http.StatusBadRequest)
			}

			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				log.Println("at upload file:", err)
			}
			defer file.Close()

			spt := strings.Split(handler.Filename, ".")
			subName := spt[len(spt)-1]

			fmt.Fprint(w, handler.Header)
			f, err := os.Create("webpage/img/" + strconv.Itoa(pdid) + "." + subName)
			if err != nil {
				log.Println("at creating file:", err)
			}
			defer f.Close()

			io.Copy(f, file)
		}
		return
	}

	path := mux.Vars(r)
	args := r.URL.Query()

	switch path["key"] {
	case "help":
		fmt.Fprint(w, ProductAPI)

	case "get":
		pdid := args.Get("pdid")

		res, err := ser.Pd.GetProductInfo(pdid)
		response(w, res, err)

	case "delete":
		val := args.Get("pdid")

		res := ser.Pd.DeleteProduct(uid, val)
		fmt.Fprint(w, res)

	case "newest":
		amount := args.Get("amount")

		res, err := ser.Pd.GetNewest(amount)
		response(w, res, err)

	case "search":
		name := args.Get("name")
		min := args.Get("minprice")
		max := args.Get("maxprice")
		eval := args.Get("eval")

		res, err := ser.Pd.SearchProducts(name, min, max, eval)
		response(w, res, err)

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
		uid, valid := ser.Ur.Login(r.FormValue("account"), r.FormValue("password"))

		if valid {
			// set session to maintain login condition
			login(w, r, uid)
			fmt.Fprintln(w, "登入成功!")
		} else {
			fmt.Fprint(w, "登入失敗")
		}

	case "regist":
		fmt.Fprint(w, ser.Ur.Regist(r.FormValue("account"), r.FormValue("password"), r.FormValue("name")))

	default:
		uid, valid := sessionValid(w, r)
		if valid {
			switch path["key"] {
			case "delete":
				fmt.Fprint(w, ser.Ur.DeleteUser(uid, r.FormValue("password")))

			case "changePassword":
				fmt.Fprint(w, ser.Ur.ChangePassword(uid, r.FormValue("oldPassword"), r.FormValue("newPassword")))

			case "changeName":
				fmt.Fprint(w, ser.Ur.ChangeName(uid, r.FormValue("newName")))

			case "logout":
				logout(w, r)
				fmt.Fprintln(w, "已登出")

			default:
				http.NotFound(w, r)
			}
		}
	}
}

func response(w http.ResponseWriter, str string, err error) {
	if err == nil {
		fmt.Fprint(w, str)
	} else {
		log.Println(err)
		http.Error(w, str, http.StatusBadRequest)
	}
}
