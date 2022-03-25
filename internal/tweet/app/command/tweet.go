package command

import "github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"

type CreateTweetHandler struct {
	TweetRepository tweet.Repository
}

func NewCreateTweetHandler(tweetRepository tweet.Repository) *CreateTweetHandler {
	return &CreateTweetHandler{TweetRepository: tweetRepository}
}

func (h CreateTweetHandler) Handle(tweet tweet.Tweet) error {
	return h.TweetRepository.Create(tweet)
}
