package search

import (
	"se/database"
)

type Search struct {
	fn *database.ProductDB
}

func (s *Search) Search(keyword string) (str []string) {

	for _, v := range s.fn.Search(keyword) {
		temp := s.fn.GetInfoFromPdID(v)
		str = append(str, temp.PdName)
	}
	return
}
