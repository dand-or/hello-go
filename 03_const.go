package main

import "fmt"

func main() {
	// 定数
	const GUNDAM = "gundam"
	const (
		ZAKU = "zaku"
		GOUF = "gouf"
		DOM  = "dom"
	)

	fmt.Println(GUNDAM, ZAKU, GOUF, DOM)
}
