package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	openFile()
}

func openFile() {
	fp, err := os.Open("sample.txt")
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(fp)
	defer fmt.Println(string(b))
	defer fp.Close()

	fmt.Println("deferは関数を抜ける前に処理を遅延実行する")
}
