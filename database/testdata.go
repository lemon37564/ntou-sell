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

	err := u.AddNewUser("test0001@gmail.com", "9279fa6a314fb0728f7cfd93669cf7f35cc01b6389fd220664919f455b307203", "測試人員A")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test0002@gmail.com", "dbc44100467a607e4653432e984eeb676302d8e070dbd3d1f66342ac0f1e7aa7", "開發人員A")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test0003@gmail.com", "75f84bcb4c96aa1f62b86ef5b2815cbc7e6cd19632c74d0a0fdf0e30a6cef297", "路人甲")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test0004@gmail.com", "dfc54fd9554e389382a8cfd4a6e69d3c4042152341422cc27c074ce1c9a313ab", "路人丁")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test0005@gmail.com", "c89d6fffc1e91b8aecce220e6fadfd49e8041f75edc5f4a7fb5e871fecca9e85", "駭客A")
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
