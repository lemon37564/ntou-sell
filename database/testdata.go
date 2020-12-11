package database

import (
	"database/sql"
	"fmt"
	"log"
)

// TestInsert tests AddNewUser with five new user
func TestInsert(db *sql.DB) {
	u := UserDBInit(db)
	p := ProductDBInit(db)

	err := u.AddNewUser("1234", "POr7EshhcPeeSrcgbKyKY3FiKPKa1HeDSdIZzts-BFo=", "測試用帳號")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test@gmail.com", "0000", "測試人員A")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test2@ntou.mail.com.tw", "ab1112c2c2", "開發人員A")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test3@gmail.com", "1234", "路人甲")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test4@gmail.com", "f06d9a3f7231c6d25724d433266f5812512ec7488c134307133e63b6b91809b7", "路人丁")
	if err != nil {
		panic(err)
	}

	err = u.AddNewUser("test5@what.com", "PASSWORD_HASH", "駭客A")
	if err != nil {
		panic(err)
	}

	_, err = p.AddNewProduct("ifone16", 2000000, "最新科技", 1, "test5@what.com", false, "null")
	if err != nil {
		panic(err)
	}

	_, err = p.AddNewProduct("ifone167", 200000000, "cioadjfdsfasdfasdbtefgdfsgfdgdsfgdsfg", 1, "test3@gmail.com", true, "2020-12-31")
	if err != nil {
		panic(err)
	}

	_, err = p.AddNewProduct("雜牌耳機", 100, "夜市貨", 16, "test@gmail.com", false, "null")
	if err != nil {
		panic(err)
	}

	_, err = p.AddNewProduct("雜牌手錶", 200, "夜市貨", 8, "test@gmail.com", false, "null")
	if err != nil {
		panic(err)
	}

	_, err = p.AddNewProduct("雜牌鞋子", 700, "夜市貨", 12, "test@gmail.com", false, "null")
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
