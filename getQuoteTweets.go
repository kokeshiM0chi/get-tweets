package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
)

const flamingId = 1145843192211787776

// とある1ツイートに対して行われている引用リツイートと、そのすべてのリプライを検索
func getQuoteTweets() {
	api := authorize()

	root, err := api.GetTweet(flamingId, url.Values{})
	if err != nil {
		fmt.Printf("error to GetTweet. err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v: %v\n", root.User.Name, root.FullText)

	allQT := qtDfs(api, root)
	// 以下で、引用リツイートに対するリプライをすべて取得する
	var s2 []anaconda.Tweet
	for _, tweet := range allQT {
		// すべての引用リツイート取得を終えたら、そのツイート1つに対してリプライを見つける
		s2 = append(s2, replyDfs(api, tweet)...)
	}

	err = writeToFile(allQT)
	if err != nil {
		fmt.Printf("Error to Json write to File. err:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("取得したツイート群をファイルに書き込みました")
}

// 引用リツイートに対する引用リツイートを探す関数
func qtDfs(api *anaconda.TwitterApi, super anaconda.Tweet) (allQT []anaconda.Tweet) {
	// 引用リツイートを検索するキーワード
	q := fmt.Sprintf("twitter.com/%v/ -from:%v", super.User.ScreenName, super.User.ScreenName)
	c := config{}
	sr, err := search(api, c, q)
	if err != nil {
		fmt.Printf("error to Search. err: %v\n", err)
	}

	var subs []anaconda.Tweet
	for _, s := range sr.Statuses {
		if s.InReplyToStatusID == super.Id {
			subs = append(subs, s)
		}
	}
	if len(subs) != 0 {
		allQT = append(allQT, subs...)
	} else {
		return nil
	}

	for _, sub := range subs {
		allQT = append(allQT, qtDfs(api, sub)...)
	}
	return allQT
}
