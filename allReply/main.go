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
		fmt.Println("検索キーワードを1つ付け加えて再度実行してください")
		fmt.Println("e.g. ./twitter-search 12345678")
		os.Exit(1)
	}
	api := lib.Authorize()
	id, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("Parse Error. Args[1]: %v, err: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	tweets, err := getReplies(api, id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	err = lib.MkFiles(tweets)
	if err != nil {
		fmt.Printf("make file error. err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("取得したツイート群をファイルに書き込みました")
}

func getReplies(api *anaconda.TwitterApi, id int64) (tweets []anaconda.Tweet, err error) {
	v := url.Values{}
	tweet, err := api.GetTweet(id, v)
	if err != nil {
		return nil, err
	}
	q := fmt.Sprintf("to:%v", tweet.User.ScreenName)
	return lib.RemoveDuplicate(lib.ReplyDfs(api, tweet, q)), nil
}
