package search

import (
	"encoding/json"
	"se/database"
)

type Search struct {
	fn *database.ProductDB
}

func (s *Search) Search(keyword string) (str string) {

	res, err := json.Marshal(s.fn.Search(keyword))
	if err != nil {
		panic(err)
	}
	return string(res)
}
