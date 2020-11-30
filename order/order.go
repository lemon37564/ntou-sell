package order

type Order struct {
	data []byte
}

func (o Order) GetOrder() []byte {
	return o.data
}
