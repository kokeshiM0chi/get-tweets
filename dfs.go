package lib

import (
	"github.com/ChimeraCoder/anaconda"
)

func ReplyDfs(api *anaconda.TwitterApi, super anaconda.Tweet, q string) (replies []anaconda.Tweet) {
	var maxId int64 = 1
	sr, err := search(api, maxId, q)
	if err != nil {
		return nil
	}

	for _, s := range sr.Statuses {
		if s.InReplyToStatusID == super.Id {
			replies = append(replies, s)
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

// func ReplyDfs(api *anaconda.TwitterApi, super anaconda.Tweet, q string) (allReplies []anaconda.Tweet) {
// 	var maxId int64 = 1
// 	sr, err := search(api, maxId, q)
// 	if err != nil {
// 		return nil
// 	}

// 	var subs []anaconda.Tweet
// 	for _, s := range sr.Statuses {
// 		if s.InReplyToStatusID == super.Id {
// 			subs = append(subs, s)
// 		}
// 	}
// 	if len(subs) == 0 {
// 		return nil
// 	}

// 	allReplies = append(allReplies, subs...)
// 	for _, sub := range subs {
// 		allReplies = append(allReplies, ReplyDfs(api, sub, q)...)
// 	}
// 	return allReplies
// }
