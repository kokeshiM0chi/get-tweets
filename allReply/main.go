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
	id, _ := strconv.ParseInt(os.Args[1], 10, 64)
	tweets, err := getReplies(id)
	err = lib.MkFiles(tweets)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("取得したツイート群をファイルに書き込みました")
}

func getReplies(id int64) (tweets []anaconda.Tweet, err error) {
	api := lib.Authorize()
	v := url.Values{}
	tweet, err := api.GetTweet(id, v)
	if err != nil {
		return nil, err
	}
	q := fmt.Sprintf("to:%v", tweet.User.ScreenName)
	// ユーザー宛てツイートの中からそのツイートに対するすべてのリプライ検索
	tweets = lib.RemoveDuplicate(lib.ReplyDfs(api, tweet, q))
	return tweets, nil
}
