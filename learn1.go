package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// 宣言
	a, b := 1, 2

	fmt.Println(a, b)

	// 足し算
	fmt.Println("2 + 5 =", 2+5)

	// 文字列の足し算
	fmt.Println("\"abc\" + \"ABC\" =", "abc"+"ABC")

	// 配列の宣言と初期化
	arr := [...]int{0, 1, 2, 3, 4}

	// for文
	for i := range arr {
		fmt.Println(i)
	}

	// string型の変数を宣言
	var ja string = "ほげほげ"

	// 文字数を出力
	fmt.Println(ja, "len:", utf8.RuneCountInString(ja))
	fmt.Println(ja, "len:", len(ja))

	// for文
	for i := 0; i < 5; i++ {

		// switch文
		switch i {
		case 0:
			fmt.Println("0")
		case 1, 2: // カンマ区切りで複数の式(値)を記述可能
			fmt.Println("1または2")
		default:
			fmt.Println("その他")
		}
	}
}
