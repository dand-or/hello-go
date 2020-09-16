package main

import "fmt"

type Book struct {
	title string
}

func main() {
	// new は領域を動的に確保し、領域へのポインタを得る
	bookList := []*Book{}
	for i := 0; i < 10; i++ {
		book := new(Book)
		book.title = fmt.Sprintf("Title#%d", i)
		bookList = append(bookList, book)
	}
	for _, book := range bookList {
		fmt.Println(book.title)
	}
}
