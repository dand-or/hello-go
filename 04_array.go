package main

import "fmt"

func main() {
	arr := [3]string{}
	arr[0] = "zaku"
	arr[1] = "gouf"
	arr[2] = "dom"

	fmt.Println(arr[0], arr[1], arr[2])

	arr2 := [2]int{1, 999}
	fmt.Println(arr2[0], arr2[1])

	// 配列の個数省略
	arr3 := [...]uint8{3, 33, 111, 11}
	fmt.Println(arr3[0], arr3[1], arr3[2], arr3[3])

	// 個数を変更可能なものをsliceという
	s := []string{}
	s = append(s, "gundam")
	s = append(s, "gundam_mk2")
	fmt.Println(s[0], s[1])

	// lenはスライスの長さ、capはメモリ上に確保されている容量
	s2 := []int{}
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
		fmt.Println(len(s2), cap(s2))
	}

	// make(スライス型, 初期個数, 初期容量) 最初からメモリ確保する
	bufa := make([]byte, 0, 1024)
}
