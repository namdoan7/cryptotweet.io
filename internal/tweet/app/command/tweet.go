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

type UpdateTweetHandler struct {
	TweetRepository tweet.Repository
}

func NewUpdateTweetHandler(tweetRepository tweet.Repository) *UpdateTweetHandler {
	return &UpdateTweetHandler{TweetRepository: tweetRepository}
}

func (h UpdateTweetHandler) Handle(tweet tweet.Tweet) error {
	return h.TweetRepository.Update(tweet)
}
