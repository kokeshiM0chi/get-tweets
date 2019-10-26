package lib

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
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
		return sr, err
	}
	return sr, nil
}

func AllSearch(api *anaconda.TwitterApi, q string) (tweets []anaconda.Tweet, err error) {
	var maxId int64 = 1
	for {
		fmt.Println(maxId)
		sr, err := search(api, maxId, q)
		if err != nil {
			return nil, err
		}
		if len(sr.Statuses) < 99 {
			// 100件未満だと同じツイート群を何度も取得してしまうため
			fmt.Println("全て取得したため、取得を終了しました")
			tweets = append(tweets, sr.Statuses...)
			break
		}
		maxId = sr.Statuses[len(sr.Statuses)-1].Id - 1
		tweets = append(tweets, sr.Statuses...)
	}
	fmt.Printf("取得ツイート数:%d\n", len(tweets))
	fmt.Println("取得したツイートに対するリプライを取得しています")
	for _, tweet := range tweets {
		// ユーザーIDに宛てられたリプライを検索
		q := fmt.Sprintf("to:%v", tweet.User.ScreenName)
		repTweets := ReplyDfs(api, tweet, q)
		if len(repTweets) != 0 {
			tweets = append(tweets, repTweets...)
		}
	}
	fmt.Printf("リプライをすべて取得しました。取得ツイート数:%d\n", len(tweets))
	tweets = RemoveDuplicate(tweets)
	fmt.Printf("重複削除後のツイート数:%d\n", len(tweets))
	return tweets, nil
}
