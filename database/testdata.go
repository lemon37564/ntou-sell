package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func TestInsert() {
	u := UserDataInit()
	defer u.DBClose()

	err := u.AddNewUser("lll@gmail.com", "1234567891012131", "哈哈哈")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("lll2@gmail.com", "1234567891012131", "哈哈哈")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("lll3@gmail.com", "1234567891012131", "哈哈哈")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("lll4@gmail.com", "1234567891012131", "哈哈哈")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("lll5@gmail.com", "1234567891012131", "哈哈哈")
	if err != nil {
		panic(err)
	}

	log.Println("insert complete")
}

// example usage
func TestSearch() {
	fmt.Println("start searching...")
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from user;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var (
		uid      string
		account  string
		password string
		name     string
		eval     float64
	)

	for rows.Next() {
		err = rows.Scan(&uid, &account, &password, &name, &eval)
		if err != nil {
			panic(err)
		}

		fmt.Println("results:")
		fmt.Println("    uid:", uid)
		fmt.Println("    account:", account)
		fmt.Println("    password hash:", password)
		fmt.Println("    name:", name)
		fmt.Println("    eval:", eval)
		fmt.Println()
	}
}
