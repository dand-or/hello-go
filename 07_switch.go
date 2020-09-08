package main

import "fmt"

func main() {
	mode := "stop"

	switch mode {
	case "start":
		fmt.Println(true)
	case "stop":
		fallthrough // breakはgoにはない。次の処理と同じにするにはfallthroughと書く
	default:
		fmt.Println(false)
	}
}
