package main

import (
	"fmt"
	"time"
)

func main() {

	// 文字列にフォーマット
	s := fmt.Sprintf("%sのは%sです。", "シアトル", "晴れ")
	fmt.Println(s)

	t := fmt.Sprintf("%d/%dの純情な感情", 1, 3)
	fmt.Println(t)

	// ローカルのタイムゾーン情報を取得
	loc, _ := time.LoadLocation("Local")
	// Time型の値を作成
	tim := time.Date(2021, 4, 13, 12, 12, 21, 14, loc)
	// 時刻を出力
	fmt.Println(tim)

	// 現在日時の取得
	date := time.Now()

	// 現在日時の出力
	fmt.Println(date)

	// 24時間後
	add := date.Add(time.Hour*24)
	fmt.Println(add)
	// 曜日取得
	fmt.Println(date.Weekday())

}
