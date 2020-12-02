package product

type Product struct {
	price       int
	amount      int
	description string
}

func NewProduct(price, amount int, description string) Product {
	return Product{price: price, amount: amount, description: description}
}

func (p Product) Price() int {
	return p.price
}

func (p Product) Amount() int {
	return p.amount
}

func (p Product) Description() string {
	return p.description
}

func (p *Product) ChangePrice(price int) {
	p.price = price
}

func (p *Product) ChangeAmount(amount int) {
	p.amount = amount
}

func (p *Product) ChangeDescription(description string) {
	p.description = description
}
