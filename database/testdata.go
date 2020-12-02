package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreate() {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userTable := `
		CREATE TABLE user(
		id varchar(16) NOT NULL,
		account varchar(256) NOT NULL,
		password_hash varchar(64) NOT NULL,
		name varchar(256),
		eval int,             
		PRIMARY KEY(id)
	);
	`

	db.Exec(userTable)
	fmt.Println("create table success.")

	// prepared statment
	stmt, err := db.Prepare("INSERT INTO user values(?,?,?,?,?);")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec("NTOUtrade-000001", "imbrian066@gmail.com",
		"f227eba4c993394502edbe8ab95b1c31dc0c45d890fbbb82b2cb9775e9bb8e6b",
		"王大明", "1")
	if err != nil {
		panic(err)
	}
	fmt.Println("insert success.")
}

func TestSearch() {
	fmt.Println("start searching...")
	db, err := sql.Open("sqlite3", "./sqlite.db")
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
		id       string
		account  string
		password string
		name     string
		eval     int
	)

	for rows.Next() {
		err = rows.Scan(&id, &account, &password, &name, &eval)
		if err != nil {
			panic(err)
		}

		fmt.Println("results:")
		fmt.Println("    id:", id)
		fmt.Println("    account:", account)
		fmt.Println("    password hash:", password)
		fmt.Println("    name:", name)
		fmt.Println("    eval:", eval)
	}
}
