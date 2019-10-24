package lib

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

func authorize() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	return anaconda.NewTwitterApi(accessToken, accessTokenSecret)
}

var jsonDir = func() string {
	user, _ := user.Current()
	return filepath.Join(user.HomeDir, "Desktop/twitter-json_data")
}()

func mkDir(jsonDir string) error {
	user, err := user.Current()
	filepath.Join(user.HomeDir)
	_, err = os.Stat(jsonDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(jsonDir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func mkFile(tweets []anaconda.Tweet) error {
	err := mkDir(jsonDir)
	if err != nil {
		fmt.Printf("ディレクトリ作成に失敗しました. err:%s\n", err)
	}
	for i, tweet := range tweets {
		json, err := json.MarshalIndent(tweet, "", "    ")
		if err != nil {
			fmt.Printf("jsonのMarshalIndent失敗. err:%s\n", err)
			return err
		}
		t := time.Now()
		file, err := os.Create(filepath.Join(
			jsonDir, fmt.Sprintf(
				"%v%v%v-%v%v%vtimeline%d.json",
				t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), i,
			),
		))
		if err != nil {
			fmt.Printf("ファイル作成失敗. err:%s\n", err)
			return err
		}
		file.Write(json)
	}
	return nil
}

func replyDfs(api *anaconda.TwitterApi, super anaconda.Tweet) (allReplies []anaconda.Tweet) {
	// ユーザーID宛のリプライを検索
	q := fmt.Sprintf("to:%v", super.User.ScreenName)
	var maxId int64 = 1
	sr, err := search(api, maxId, q)
	var subs []anaconda.Tweet
	for _, s := range sr.Statuses {
		if s.InReplyToStatusID == super.Id {
			subs = append(subs, s)
		}
	}
	if len(subs) != 0 {
		allReplies = append(allReplies, subs...)
	} else {
		return nil
	}
	for _, sub := range subs {
		allReplies = append(allReplies, replyDfs(api, sub)...)
	}
	return allReplies
}

func removeDuplicate(tweets []anaconda.Tweet) (distinctTweets []anaconda.Tweet) {
	m := make(map[int64]struct{})
	for _, tweet := range tweets {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[tweet.Id]; !ok {
			m[tweet.Id] = struct{}{}
			distinctTweets = append(distinctTweets, tweet)
		}
	}
	return distinctTweets
}
