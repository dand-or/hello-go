package main

import "fmt"

func main() {
	var a1 int
	var p1 *int

	p1 = &a1
	*p1 = 1000
	fmt.Println(a1)

	x1 := 123
	x2 := 123
	hoge(x1, &x2)
	fmt.Println(x1, x2)
}

// b1は値渡し　b2はポインタ渡し
func hoge(b1 int, b2 *int) {
	b1 = 456
	*b2 = 456
}
