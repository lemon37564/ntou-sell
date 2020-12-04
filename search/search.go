package search

import (
	"encoding/json"
	"se/database")


type Search struct {
	fn *database.ProductDB
}

func (s *Search) Search(keyword string) (str string) {
	var prods := []database.Product
	p
	for _, v := range s.fn.Search(keyword) {
		temp := s.fn.GetInfoFromPdID(v)
		
	}

	
}
