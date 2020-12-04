package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TestInsert tests AddNewUser with five new user
func TestInsert(db *sql.DB) {
	u := UserDBInit(db)

	err := u.AddNewUser("test0001@gmail.com", "password_hash", "測試人員A")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test0002@ntou.mail.com.tw", "ab1112c2c2", "開發人員A")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test0003@yahoo.com.tw", "ef297", "路人甲")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test0004@gmail.com", "e04e04e04", "路人丁")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test0005@what.com", "jjj09090", "駭客A")
	if err != nil {
		panic(err)
	}

	log.Println("insert complete")
}

// TestSearch shows all the users
func TestSearch(db *sql.DB) {
	fmt.Println("start searching...")

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
