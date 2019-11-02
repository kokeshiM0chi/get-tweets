package main

import (
	"fmt"
	"getTweets"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("検索キーワードを1つ付け加えて再度実行してください.\n e.g. ./twitter-search 12345678\n")
		os.Exit(1)
	}
	id, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("Parse Error. Args[1]: %v, err: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	tweets, err := getQuote(lib.Authorize(), id)
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

// 特定ツイートの中で、引用リツイートされたもののみを取得する
func getQuote(api *anaconda.TwitterApi, id int64) (tweets []anaconda.Tweet, err error) {
	v := url.Values{}
	tweet, err := api.GetTweet(id, v)
	if err != nil {
		return nil, err
	}
	// 引用リツイート検索
	tweets, err = lib.AllSearch(api, fmt.Sprintf(
		"twitter.com/%v/ -from:%v", tweet.User.ScreenName, tweet.User.ScreenName), true)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
