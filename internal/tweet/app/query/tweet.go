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

type GetTweetHandler struct {
	TweetRepository tweet.Repository
}

func NewGetTweetHandler(tweetRepository tweet.Repository) *GetTweetHandler {
	return &GetTweetHandler{TweetRepository: tweetRepository}
}

func (h GetTweetHandler) Handle(tweetId string) (*tweet.Tweet, error) {
	return h.TweetRepository.GetTweet(tweetId)
}
