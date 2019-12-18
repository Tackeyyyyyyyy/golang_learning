package main

import "fmt"

// 掛け算
func multiplication(a int, b int) int {

	// 戻り値を返す
	return a * b
}

// 構造体型の宣言
type Employee struct {
	name string
	age  int
}

type Company struct {
	id int
	Employee
}

// main関数はパラメータ、戻り値ともに持たない
func main() {

	// multiplication関数を呼び出す
	answer := multiplication(10, 2)

	// 関数から返された戻り値を出力
	fmt.Println(answer) // 20

	// 構造体リテラルで初期化
	e := Company{1, Employee{"長澤まさみ", 32}}

	// 出力
	fmt.Println(e)

	// ポインタも構造体リテラルで作成可能
	p4 := &Employee{"橋本環奈", 36}

	// 出力
	fmt.Println(p4)

	// 構造体型の変数を宣言
	var x struct {
		i1, i2 int     // int型フィールド
		f      float32 // float32型フィールド
		s      string  // string型フィールド
	}

	// 構造体の各フィールドに値をセット
	x.i1 = 12
	x.i2 = 213
	x.f = 3.321
	x.s = "line * yahoo"

	// 出力
	fmt.Println(x)
}
