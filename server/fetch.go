package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	// temp area
	if r.Method == "POST" {
		log.Println("receive post (product)")
		if mux.Vars(r)["key"] == "postadd" {
			var pdid int
			var stat error

			r.ParseMultipartForm(32 << 20)

			name := r.Form["name"][0]
			price := r.Form["price"][0]
			des := r.Form["description"][0]
			amount := r.Form["amount"][0]
			bid := r.Form["bid"][0]
			date := r.Form["date"][0]

			_, err := ser.Pd.AddProduct(uid, name, price, des, amount, bid, date)
			if err == nil {
				fmt.Fprint(w, "ok")
			} else {
				http.Error(w, "failed", http.StatusBadRequest)
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
		name := args.Get("name")
		price := args.Get("price")
		des := args.Get("description")
		amount := args.Get("amount")
		bid := args.Get("bid")
		date := args.Get("date")

		_, err := ser.Pd.AddProduct(uid, name, price, des, amount, bid, date)
		if err == nil {
			fmt.Fprint(w, "ok")
		} else {
			http.Error(w, "failed", http.StatusBadRequest)
		}

	case "get":
		pdid := args.Get("pdid")

		res, err := ser.Pd.GetProductInfo(pdid)
		response(w, res, err)

	case "delete":
		val := args.Get("pdname")

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

func response(w http.ResponseWriter, str string, err error) {
	if err == nil {
		fmt.Fprint(w, str)
	} else {
		log.Println(err)
		http.Error(w, str, http.StatusBadRequest)
	}
}
