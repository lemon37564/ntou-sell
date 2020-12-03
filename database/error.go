package database

import "fmt"

type HashValError struct {
	length int
}

func (he HashValError) Error() string {
	return "hash value length unexpected: " + fmt.Sprintf("%v\n", he.length)
}
