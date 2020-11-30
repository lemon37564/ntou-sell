package cart

type cart struct {
	data []byte
}

func (c cart) GetData() []byte {
	return c.data
}
