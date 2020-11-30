package product

type Product struct {
	data []byte
}

func (p Product) GetInfo() []byte {
	return p.data
}
