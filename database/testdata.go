package database

import (
	"database/sql"
	"log"
	"time"
)

// TestInsert tests AddNewUser with five new user
func TestInsert(db *sql.DB) {
	u := UserDBInit(db)
	p := ProductDBInit(db)

	err := u.AddNewUser("1234", "DinaSXJlIqL7-PmiEJBJmbhijzeJhSHiqyD5Jx5S1D0=", "測試用帳號")
	if err != nil {
		log.Println(err)
	}

	err = u.AddNewUser("abcd", "DinaSXJlIqL7-PmiEJBJmbhijzeJhSHiqyD5Jx5S1D0=", "除錯人員ABC")
	if err != nil {
		log.Println(err)
	}

	err = u.AddNewUser("test2@ntou.mail.com.tw", "25724d43266f5ab1125724d4332f52c2c2", "開發人員A")
	if err != nil {
		log.Println(err)
	}

	err = u.AddNewUser("test3@gmail.com", "9a3f7231c6d25724d433266f5812513e63b6b", "路人甲")
	if err != nil {
		log.Println(err)
	}

	err = u.AddNewUser("test4@gmail.com", "f06d9a3f7231c6d25724d433266f5812512ec7488c134307133e63b6b91809b7", "路人丁")
	if err != nil {
		log.Println(err)
	}

	err = u.AddNewUser("test5@what.com", "PA33e63b6bSSWf06d9a3f7231c6d25724d433266f5812512ecH", "駭客A")
	if err != nil {
		log.Println(err)
	}

	_, err = p.AddNewProduct("ifone16", 2000000, "最新科技", 1, 2, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = p.AddNewProduct("ifone167", 200000000, "cioadjfdsfasdfasdbtefgdfsgfdgdsfgdsfg", 1, 1, true, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = p.AddNewProduct("雜牌耳機", 100, "夜市貨", 16, 2, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = p.AddNewProduct("雜牌手錶", 200, "夜市貨", 8, 3, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = p.AddNewProduct("雜牌鞋子", 700, "夜市貨", 12, 1, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	log.Println("Test insert complete")
}

// TestSearch shows all the users
// func TestSearch(db *sql.DB) {
// 	fmt.Println("start searching...")

// 	rows, err := db.Query("select * from user;")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	var (
// 		uid      string
// 		account  string
// 		password string
// 		name     string
// 		eval     float64
// 	)

// 	for rows.Next() {
// 		err = rows.Scan(&uid, &account, &password, &name, &eval)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("results:")
// 		fmt.Println("    uid:", uid)
// 		fmt.Println("    account:", account)
// 		fmt.Println("    password hash:", password)
// 		fmt.Println("    name:", name)
// 		fmt.Println("    eval:", eval)
// 		fmt.Println()
// 	}
// }
