package database

import "os"

const (
	root = "C:/software-engineering"
	name = "database.db"
	file = root + "/" + name
)

// check if there's database exists
// if no, init.
func CheckDB() {
	_, err := os.Stat(root)
	if err != nil {
		createFile()
	}

	_, err = os.Stat(file)
	if err != nil {
		createTables()
	}
}

func createFile() {
	err := os.Mkdir(root, 0666)
	if err != nil {
		panic(err)
	}
}

// FATAL: this command will remove whole database
func RemoveAll() {
	err := os.RemoveAll(root)
	if err != nil {
		panic(err)
	}
}
