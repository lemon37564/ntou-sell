package cart

import (
	"se/database"
)

type cart struct {
	db database.CartData
}

// // AddProduct return true if add success
// func (c *cart) AddProduct(p product.Product, amount int) bool {
// 	c.db.Insert()
// 	return true
// }

// // RemoveProduct remove product in the cart if exists. if there's no such product in the cart, return false
// func (c *cart) RemoveProduct(p product.Product) bool {
// 	c.db.Delete()
// 	return true
// }

// // ModifyAmount changes the amount of specific product. returns ture if success
// func (c *cart) ModifyAmount(p product.Product, newAmount int) bool {
// 	c.db.UpdateAmount()
// 	return true
// }

// // GetProducts returns all the product in cart
// func (c cart) GetProducts() map[product.Product]int {
// 	c.db.Select()
// }

// // TotalCount returns how many different products in the cart
// func (c cart) TotalCount() int {
// 	c.db.Select()
// }

// // Sum returns the total price of products in the cart
// func (c cart) Sum() (sum int) {
// 	for i, v := range c.products {
// 		sum += i.Price() * v
// 	}

// 	return
// }
