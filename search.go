package lib

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"strconv"
)

func search(api *anaconda.TwitterApi, maxId int64, q string) (sr anaconda.SearchResponse, err error) {
	v := url.Values{}
	v.Set("count", "100")
	if maxId != 0 {
		v.Set("max_id", strconv.FormatInt(maxId, 10))
	}
	sr, err = api.GetSearch(q, v)
	if err != nil {
		print(err)
		return sr, err
	}
	return sr, nil
}

func AllSearch(q string) (tweets []anaconda.Tweet) {
	api := Authorize()
	var maxId int64 = 1
	for {
		sr, err := search(api, maxId, q)
		if err != nil {
			fmt.Printf("検索失敗. err:%v\n", err)
			os.Exit(1)
		}
		if len(sr.Statuses) < 99 {
			// 100件未満だと同じツイート群を何度も取得してしまうため
			fmt.Println("全て取得したため、取得を終了しました")
			tweets = append(tweets, sr.Statuses...)
			break
		}
		maxId = sr.Statuses[len(sr.Statuses)-1].Id - 1 //statuses末尾取得
		tweets = append(tweets, sr.Statuses...)
	}
	var d []anaconda.Tweet
	fmt.Printf("取得ツイート数:%d\n", len(tweets))
	fmt.Println("取得したツイートに対するリプライを取得しています")
	for _, tweet := range tweets {
		// ユーザーID宛のリプライを検索
		// q := fmt.Sprintf("to:%v", super.User.ScreenName)
		d = ReplyDfs(api, tweet, q)
	}
	fmt.Println("リプライをすべて取得しました")
	fmt.Printf("取得ツイート数:%d\n", len(tweets))
	tweets = append(tweets, d...)
	return tweets
}
