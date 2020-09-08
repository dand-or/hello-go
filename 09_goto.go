package main

import "fmt"

func main() {
	// goにはtry catchがない
	err := true
	if err {
		goto Done
	}
	fmt.Println("hoge")

Done:
	fmt.Println(err)
}
