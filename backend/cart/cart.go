package cart

import (
	"database/sql"
	"encoding/json"
	"log"
	"se/database"
	"strconv"
)

type Cart struct {
	db  *database.CartDB
	db2 *database.ProductDB
}

func NewCart(db *sql.DB) *Cart {
	c := new(Cart)
	c.db = database.CartDBInit(db)

	return c

}

// AddProduct return true if add success
func (c *Cart) AddProductToCart(id, pdid, amount int) bool {
	err := c.db.AddProductIntoCart(id, pdid, amount)
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
	pds := c.db.GetAllProductOfUser(uid)

	res, err := json.Marshal(pds)
	if err != nil {

		return "Something Wrong in Getting Data"
	}

	return string(res)
}

// TotalCount returns how many different products in the cart
func (c *Cart) TotalCount(uid int) string {
	if c.db.Total == 0 {
		c.GetProducts(uid)
	}

	return strconv.Itoa(c.db.Total)
}

// Sum returns the total price of products in the cart
// func (c Cart) Sum() (sum int) {
// 	for i, v := range c.products {
// 		sum += i.Price() * v
// 	}

// 	return
// }
