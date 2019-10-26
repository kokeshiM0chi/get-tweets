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

func Authorize() *anaconda.TwitterApi {
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

func MkFiles(tweets []anaconda.Tweet) (err error) {
	if err = mkDir(jsonDir); err != nil {
		return err
	}
	for i, tweet := range tweets {
		err = func() error {
			json, err := json.MarshalIndent(tweet, "", "    ")
			if err != nil {
				return err
			}
			t := time.Now()
			file, err := os.Create(filepath.Join(
				jsonDir, fmt.Sprintf(
					"%v%v%v-%v%vtweet%d.json",
					t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), i,
				),
			))
			if err != nil {
				return err
			}
			file.Write(json)
			defer file.Close()
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveDuplicate(tweets []anaconda.Tweet) (distinctTweets []anaconda.Tweet) {
	m := make(map[int64]struct{})
	for _, tweet := range tweets {
		if _, ok := m[tweet.Id]; !ok {
			m[tweet.Id] = struct{}{}
			distinctTweets = append(distinctTweets, tweet)
		}
	}
	return distinctTweets
}

// func MkFile(tweet anaconda.Tweet, i int64) error {
// 	if err := mkDir(jsonDir); err != nil {
// 		fmt.Printf("ディレクトリ作成に失敗しました. err:%s\n", err)
// 	}
// 	json, err := json.MarshalIndent(tweet, "", "    ")
// 	if err != nil {
// 		fmt.Printf("jsonのMarshalIndent失敗. err:%s\n", err)
// 		return err
// 	}
// 	t := time.Now()
// 	file, err := os.Create(filepath.Join(
// 		jsonDir, fmt.Sprintf(
// 			"%v%v%v-%v%vtimeline%d.json",
// 			t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), i,
// 		),
// 	))
// 	if err != nil {
// 		fmt.Printf("ファイル作成失敗. err:%s\n", err)
// 		return err
// 	}
// 	file.Write(json)
// 	defer file.Close()
// 	return nil
// }
