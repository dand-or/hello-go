package main

import "fmt"

func main() {
	x := 0
	y := 3
	// goにwhileはなくループは全てfor
	for x < y {
		x++
		fmt.Println(x)
	}

	// いつものfor
	for n := 0; n < 5; n++ {
		fmt.Println(n)
	}

	// 無限ループ
	m := 0
	for {
		m++
		if m > 10 {
			break
		} else if m%2 == 1 {
			continue
		} else {
			fmt.Println(m)
		}
	}

	// 配列やスライスなどiterateなものにはrangeでforeach的なループ
	colors := [...]string{"zaku", "gouf", "dom"}
	for i, color := range colors {
		fmt.Printf("%d: %s\n", i, color)
	}
}
