package main

import (
	"fmt"
	"time"
)

func funcA() {
	for i := 0; i < 10; i++ {
		fmt.Print("A")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go funcA()
	for i := 0; i < 10; i++ {
		fmt.Print("M")
		time.Sleep(20 * time.Millisecond)
	}
}
