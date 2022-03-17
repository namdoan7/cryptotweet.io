package command

import "github.com/cryptotweet.io/internal/tweet/domain/tweet"

type CreateTweetHandler struct {
	TweetRepository tweet.Repository
}

func NewCreateTweetHandler(tweetRepo tweet.Repository) *CreateTweetHandler {
	return &CreateTweetHandler{TweetRepository: tweetRepo}
}

func (h CreateTweetHandler) Handle(tweet tweet.Tweet) error {
	return h.TweetRepository.Create(tweet)
}
