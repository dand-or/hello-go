package main

import "fmt"

type Printable interface {
	ToString() string
}

func PrintOut(p Printable) {
	fmt.Println(p.ToString())
}

type Gundam struct {
	name string
}

func (g Gundam) ToString() string {
	return g.name
}

type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}

func main() {
	g := Gundam{"zeta_gundam"}
	b := Book{"ホビージャパン"}
	PrintOut(g)
	PrintOut(b)
}
