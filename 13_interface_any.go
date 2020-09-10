package main

import "fmt"

// 関数がないインターフェイスinterface{}はany型のように扱える
func hoge(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Printf("%t\n", a)
	case int:
		fmt.Printf("%d\n", a)
	case string:
		fmt.Printf("%s\n", a)
	}
}

func main() {
	hoge(true)
	hoge(2220)
	hoge("foo")
}
