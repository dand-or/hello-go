package main

import "fmt"

func main() {
	// var 変数名 型
	var x1 int

	// 初期値セット
	var x2 int = 200

	// 初期値セット 型が明確
	var x3 = 300

	// := でvarを省略
	x4 := 400

	// まとめて変数定義
	var (
		x5 = 500
		x6 = 600
	)

	// 代入
	x1 = 100

	// 複数の値の同時代入
	var name string
	var age int
	name, age = "pikachu", 1

	fmt.Println(x1, x2, x3, x4, x5, x6, name, age)
}
