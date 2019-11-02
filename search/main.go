package main

import (
	"fmt"
	"os"
	// "github.com/ssabcire/get-tweets"
	"getTweets"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("検索キーワードを1つ付け加えて再度実行してください.\n e.g. ./twitter-search abc\n")
		os.Exit(1)
	}
	api := lib.Authorize()
	q := os.Args[1]
	tweets, err := lib.AllSearch(api, q, true)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	err = lib.MkFiles(tweets)
	if err != nil {
		fmt.Printf("make file error. err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("取得したツイート群をファイルに書き込みました")
}
