package cart

import "se/product"

type Cart struct {
	// key->product, value->amount
	products map[product.Product]int
}

// AddProduct return true if add success
func (c *Cart) AddProduct(p product.Product, amount int) bool {
	_, exist := c.products[p]
	if exist {
		return false
	}

	c.products[p] = amount
	return true
}

// RemoveProduct remove product in the cart if exists. if there's no such product in the cart, return false
func (c *Cart) RemoveProduct(p product.Product) bool {
	_, exist := c.products[p]
	if !exist {
		return false
	}

	delete(c.products, p)
	return true
}

// ModifyAmount changes the amount of specific product. returns ture if success
func (c *Cart) ModifyAmount(p product.Product, newAmount int) bool {
	_, exists := c.products[p]
	if !exists {
		return false
	}

	c.products[p] = newAmount
	return true
}

// GetProducts returns all the product in cart
func (c Cart) GetProducts() map[product.Product]int {
	return c.products
}

// TotalCount returns how many different products in the cart
func (c Cart) TotalCount() int {
	return len(c.products)
}

// Sum returns the total price of products in the cart
func (c Cart) Sum() (sum int) {
	for i, v := range c.products {
		sum += i.Price() * v
	}

	return
}
