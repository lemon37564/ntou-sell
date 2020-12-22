package database

import (
	"log"
	"time"
)

// TestInsert tests AddNewUser and AddNewProduct
func TestInsert() {
	dt := OpenAndInit()

	err := dt.AddNewUser("1234", "DinaSXJlIqL7-PmiEJBJmbhijzeJhSHiqyD5Jx5S1D0=", "測試用帳號")
	if err != nil {
		log.Println(err)
	}

	err = dt.AddNewUser("abcd", "DinaSXJlIqL7-PmiEJBJmbhijzeJhSHiqyD5Jx5S1D0=", "除錯人員ABC")
	if err != nil {
		log.Println(err)
	}

	_, err = dt.AddProduct("ifone16", 2000000, "最新科技", 1, 2, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = dt.AddProduct("ifone167", 200000000, "cioadjfdsfasdfasdbtefgdfsgfdgdsfgdsfg", 1, 1, true, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = dt.AddProduct("雜牌耳機", 100, "夜市貨", 16, 2, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = dt.AddProduct("雜牌手錶", 200, "夜市貨", 8, 3, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = dt.AddProduct("雜牌鞋子", 700, "夜市貨", 12, 1, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	log.Println("Test insert complete")
}
