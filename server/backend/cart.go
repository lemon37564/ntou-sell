package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"se/database"
)

// Cart is a module that handle cart functions
type Cart struct {
	db *database.CartDB
}

// CartInit return cart module
func CartInit(db *sql.DB) *Cart {
	c := new(Cart)
	c.db = database.CartDBInit(db)

	return c

}

// AddProductToCart return true if add success
func (c *Cart) AddProductToCart(uid, pdid, amount int) bool {
	err := c.db.AddProductIntoCart(uid, pdid, amount)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// RemoveProduct remove product in the cart if exists. if there's no such product in the cart, return false
func (c *Cart) RemoveProduct(id, pdid int) string {
	err := c.db.DeleteProductFromCart(id, pdid)
	if err != nil {
		log.Println(err)
		return "fail to remove from cart"
	}
	return "Success"
}

// ModifyAmount changes the amount of specific product. returns ture if success
func (c *Cart) ModifyAmount(uid, pdid, amount int) bool {
	err := c.db.UpdateAmount(uid, pdid, amount)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// GetProducts returns all the product in cart
func (c Cart) GetProducts(uid int) string {
	pds, _ := c.db.GetAllProductOfUser(uid)

	res, err := json.Marshal(pds)
	if err != nil {

		return "Something Wrong in Getting Data"
	}

	return string(res)
}
