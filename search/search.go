package search

import (
	"encoding/json"
	"log"
	"se/database"
)

type Search struct {
	fn *database.ProductDB
}

func (s *Search) Search(keyword string) (str string) {

	res, err := json.Marshal(s.fn.Search(keyword))
	if err != nil {
		log.Println(err)
		return "Fail"
	}
	return string(res)
}
