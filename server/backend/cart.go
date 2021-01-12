package backend

import (
	"encoding/json"
	"log"
	"se/database"
	"strconv"
)

// Cart is a module that handle cart functions
type Cart struct {
	fn *database.Data
}

// CartInit return cart module
func CartInit(data *database.Data) *Cart {
	return &Cart{fn: data}
}

// AddProductToCart return true if add success
func (c *Cart) AddProductToCart(uid int, rawPdid, rawAmount string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return "cannot convert " + rawAmount + " into integer", err
	}

	err = c.fn.AddProductIntoCart(uid, pdid, amount)
	if err != nil {
		log.Println(err)
		return "false", err
	}
	return "true", nil
}

// RemoveProduct remove product in the cart if exists. if there's no such product in the cart, return false
func (c *Cart) RemoveProduct(uid int, rawPdid string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	err = c.fn.DeleteProductFromCart(uid, pdid)
	if err != nil {
		log.Println(err)
		return "fail to remove from cart", err
	}
	return "Success", nil
}

// ModifyAmount changes the amount of specific product. returns ture if success
func (c *Cart) ModifyAmount(uid int, rawPdid, rawAmount string) (string, error) {
	pdid, err := strconv.Atoi(rawPdid)
	if err != nil {
		return "cannot convert " + rawPdid + " into integer", err
	}

	amount, err := strconv.Atoi(rawAmount)
	if err != nil {
		return "cannot convert " + rawAmount + " into integer", err
	}

	err = c.fn.UpdateCartAmount(uid, pdid, amount)
	if err != nil {
		log.Println(err)
		return "false", err
	}
	return "true", err
}

// GetProducts returns all the product in cart
func (c Cart) GetProducts(uid int) string {
	pds, _ := c.fn.GetAllProductOfUser(uid)

	res, err := json.Marshal(pds)
	if err != nil {

		return "Something Wrong in Getting Data"
	}

	return string(res)
}
