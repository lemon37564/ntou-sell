package history

type History struct {
	data []byte
}

func (h History) GetHistory() []byte {
	return h.data
}
