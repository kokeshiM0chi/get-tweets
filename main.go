package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("検索キーワードを1つ付け加えて再度実行してください")
		fmt.Println("e.g. ./twitter-search abc")
		os.Exit(1)
	}
	q := os.Args[1]
	allSearch(q)
}
