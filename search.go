package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
)

func search(api *anaconda.TwitterApi, maxIdStr string, q string) (sr anaconda.SearchResponse, err error) {
	v := url.Values{}
	v.Set("count", "100")
	if maxIdStr != "" {
		fmt.Println(maxIdStr)
		v.Set("max_id", maxIdStr)
	}
	sr, err = api.GetSearch(q, v)
	if err != nil {
		print(err)
		return sr, err
	}
	return sr, nil
}

func allSearch(q string) {
	api := authorize()
	maxIdStr := "-1"
	var s []anaconda.Tweet
	for {
		sr, err := search(api, maxIdStr, q)
		if err != nil {
			fmt.Printf(" err:%v\n", err)
			os.Exit(1)
		}
		if len(sr.Statuses) < 99 {
			// 100件未満だと同じツイート群を何度も取得してしまうため
			fmt.Println("全て取得したため、取得を終了しました")
			s = append(s, sr.Statuses...)
			break
		}
		maxIdStr = sr.Statuses[len(sr.Statuses)-1].IdStr //statuses末尾取得
		s = append(s, sr.Statuses...)
	}
	// var d []anaconda.Tweet
	// for _, tweet := range s {
	// 	d = replyDfs(api, tweet)
	// }
	// s = append(s, d...)

	err := mkFile(s)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("取得したツイート群をファイルに書き込みました")
}
