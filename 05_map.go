package main

import "fmt"

func main() {
	// hashとか、C#でいうとこのDictionary

	// map[キーの型]バリューの型
	m1 := map[string]int{
		"science": 30,
		"english": 44,
	}
	fmt.Println(m1["science"], m1["english"])

	// 要素追加
	m1["math"] = 2

	// 要素削除
	delete(m1, "english")

	// マップ個数
	fmt.Println(len(m1))

	// マップに要素が存在するかわかる
	_, ok := m1["english"]
	if ok {
		fmt.Println("exist")
	} else {
		fmt.Println("nothing")
	}

	// マップをループ
	for key, value := range m1 {
		fmt.Printf("%s=%d\n", key, value)
	}
}
