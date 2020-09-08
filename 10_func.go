package main

import "fmt"

func main() {
	fmt.Println(isEven(1000))
	fmt.Println(multiReturn(2, 5))

	// 複数地の返却で不要なものは_で捨てれる
	_, a := multiReturn(3, 4)
	fmt.Println(a)

	kahen("test", 1, 2, 3, 4, 5, 6, 7)
}

func isEven(x int) bool {
	return x%2 == 0
}

// 複数値の返却
func multiReturn(x int, y int) (int, int) {
	return x + y, x * y
}

// 可変引数
func kahen(a string, b ...int) {
	for i, item := range b {
		fmt.Println(a, i, item)
	}
}
