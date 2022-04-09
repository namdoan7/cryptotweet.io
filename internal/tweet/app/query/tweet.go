package query

import "github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"

type ListTweetsHandler struct {
	TweetRepository tweet.Repository
}

func NewListTweetsHandler(tweetRepository tweet.Repository) *ListTweetsHandler {
	return &ListTweetsHandler{TweetRepository: tweetRepository}
}

func (h ListTweetsHandler) Handle() ([]tweet.Tweet, error) {
	return h.TweetRepository.Find()
}
