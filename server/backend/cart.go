package backend

import (
	"encoding/json"
	"log"
	"se/database"
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
func (c *Cart) AddProductToCart(uid, pdid, amount int) bool {
	err := c.fn.AddProductIntoCart(uid, pdid, amount)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// RemoveProduct remove product in the cart if exists. if there's no such product in the cart, return false
func (c *Cart) RemoveProduct(id, pdid int) string {
	err := c.fn.DeleteProductFromCart(id, pdid)
	if err != nil {
		log.Println(err)
		return "fail to remove from cart"
	}
	return "Success"
}

// ModifyAmount changes the amount of specific product. returns ture if success
func (c *Cart) ModifyAmount(uid, pdid, amount int) bool {
	err := c.fn.UpdateCartAmount(uid, pdid, amount)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
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
