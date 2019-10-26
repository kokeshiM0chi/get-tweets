package lib

import (
	"github.com/ChimeraCoder/anaconda"
)

func ReplyDfs(api *anaconda.TwitterApi, super anaconda.Tweet, q string) (replies []anaconda.Tweet) {
	var maxId int64 = 1
	var tweets []anaconda.Tweet
	for {
		sr, err := search(api, maxId, q)
		if err != nil {
			return nil
		}
		if len(sr.Statuses) < 99 {
			// 100件未満だと同じツイート群を何度も取得してしまうため
			tweets = append(tweets, sr.Statuses...)
			break
		}
		maxId = sr.Statuses[len(sr.Statuses)-1].Id - 1
		tweets = append(tweets, sr.Statuses...)
	}

	for _, tweet := range tweets {
		if tweet.InReplyToStatusID == super.Id {
			replies = append(replies, tweet)
		}
	}
	if len(replies) == 0 {
		return nil
	}

	for _, sub := range replies {
		replies = append(replies, ReplyDfs(api, sub, q)...)
	}
	return replies
}
