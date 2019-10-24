package lib

import (
	// "fmt"
	"github.com/ChimeraCoder/anaconda"
)

func ReplyDfs(api *anaconda.TwitterApi, super anaconda.Tweet, q string) (allReplies []anaconda.Tweet) {
	var maxId int64 = 1
	sr, err := search(api, maxId, q)
	if err != nil {
	}
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
		allReplies = append(allReplies, ReplyDfs(api, sub, q)...)
	}
	return allReplies
}
