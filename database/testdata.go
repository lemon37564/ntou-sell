package database

import (
	"log"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// TestInsert tests AddNewUser and AddNewProduct
func TestInsert() {
	return

	rand.Seed(time.Now().UnixNano())

	err := AddNewUser(RandStringRunes(6), "DinaSXJlIqL7-PJx5S1D0=", "測試用帳號")
	if err != nil {
		panic(err)
	}

	err = AddNewUser(RandStringRunes(6), "DinaSjzeJhHiqyD5Jx5S1D0=", "除錯人員ABC")
	if err != nil {
		panic(err)
	}

	_, err = AddProduct("ifone16", 200000, "最新科技", 1, 2, false, time.Now())
	if err != nil {
		panic(err)
	}

	_, err = AddProduct("ifone167", 20000000, "來自未來的產物", 1, 1, true, time.Now())
	if err != nil {
		panic(err)
	}

	_, err = AddProduct("ifone3310", 987654321, "沁・ｷ@ﾐwｻﾛBASﾀ・tlｼｼﾈJRrｒK-'炭ﾋ函mﾈ妲JZ3M莇ﾉﾘ4ﾀ・q･ｱ･｡U*・縣ｦ・J・k3ﾙﾆXd2KﾝM0%配・朽ｧ从aDｱ#・*ｽy啼ｪﾙﾕﾊIｾ・lﾀUlZ呎4扈n歪晞德f2ｧh#董_]Aｼrｱ｣黔W殉Rﾝ>・ﾍﾔ[ｻｻｮ・賻ｨq(ｽｼｸ較ﾒｌCﾈfﾂA・ﾃ劜&B・ ｩ%ﾆﾆl[粲ﾘp'ｳ｢ｪmwｱ・U5LQ｢;]N8・t6蒹}・Qﾌﾋﾅｿl,Pｶ赶摶+lpWﾋYﾆn	(+ﾚuq6ﾋ[Vｰ囀ｬｹ塒ﾉm5ﾃ箆濬3=・6ﾂ慶~渺ｴﾅﾗ・HlW气ﾔｷｽ孝ﾐtMﾀJｸD;白", 1, 1, true, time.Now())
	if err != nil {
		panic(err)
	}

	_, err = AddProduct("雜牌耳機", 100, "夜市貨", 16, 1, false, time.Now())
	if err != nil {
		panic(err)
	}

	_, err = AddProduct("雜牌手錶", 200, "夜市貨", 8, 1, false, time.Now())
	if err != nil {
		panic(err)
	}

	_, err = AddProduct("雜牌鞋子", 700, "夜市貨", 12, 2, false, time.Now())
	if err != nil {
		panic(err)
	}

	log.Println("Test insert complete")
}
