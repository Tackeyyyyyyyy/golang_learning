package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// 配列を宣言
	x := [5]string{"1", "2", "3", "4", "5"}

	// スライス型の変数を宣言
	var s1 []string

	// 配列全体をスライス
	s1 = x[:]
	fmt.Println(s1)

	// ファイルのオープン
	file, err := os.Open("test.text")

	// オープンに失敗したか判定
	if err != nil {

		// エラーの詳細情報を出力
		fmt.Println(err.Error())

		// 終了
		os.Exit(1)
	}
	// ファイルの名前
	fmt.Println(file.Name())
	// ファイルのクローズ

	// ioutilを使った方
	data, err := ioutil.ReadFile("test.text")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
