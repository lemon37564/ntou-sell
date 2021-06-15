package backend

import (
	"encoding/json"
	"log"
	"se/database"
	"strconv"
)

// AddProductToCart return true if add success
func AddProductToCart(uid int, rawPdid, rawAmount string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return "cannot convert " + rawAmount + " into integer", err
	}

	err = database.AddProductIntoCart(uid, pdid, amount)
	if err != nil {
		log.Println(err)
		return "false", err
	}
	return "true", nil
}

// RemoveProduct remove product in the cart if exists. if there's no such product in the cart, return false
func RemoveProduct(uid int, rawPdid string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	err = database.DeleteProductFromCart(uid, pdid)
	if err != nil {
		log.Println(err)
		return "fail to remove from cart", err
	}
	return "Success", nil
}

// ModifyAmount changes the amount of specific product. returns ture if success
func ModifyProductAmount(uid int, rawPdid, rawAmount string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return "cannot convert " + rawAmount + " into integer", err
	}

	err = database.UpdateCartAmount(uid, pdid, amount)
	if err != nil {
		log.Println(err)
		return "false", err
	}
	return "true", err
}

// GetProducts returns all the product in cart
func GetProducts(uid int) string {
	pds, _ := database.GetAllProductOfUser(uid)

	res, err := json.Marshal(pds)
	if err != nil {

		return "Something Wrong in Getting Data"
	}

	return string(res)
}
