package main

import "fmt"

// goではclassの代わりにstructを使う
type Person struct {
	Name string // 大文字で始まるとpublic
	age  int    // 小文字で始まるとprivate
}

// structにはメンバー変数のみ定義してメソッドは{this相当の変数 *struct名}をつけたfuncを書く
func (p *Person) SetPerson(name string, age int) {
	p.Name = name
	p.age = age
}

func (p *Person) GetAge() int {
	return p.age
}

func main() {
	var p1 Person
	p1.SetPerson("gundam", 2)
	fmt.Println(p1.Name, p1.GetAge())

	p2 := Person{"zaku", 199}
	fmt.Println(p2)
}
