package search

import "se/database"

type Search struct {
	fn *database.ProductDB
}

func (s *Search) Search(keyword string) []int {
	return s.fn.Search()
}
