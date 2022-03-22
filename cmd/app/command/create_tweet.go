package command

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
)

type CreateTweet struct {
	Tweet *tweet.Tweet
}

type CreateTweetHandler struct {
	TweetService TweetService
}

func NewCreateTweetHandler(tweetService TweetService) CreateTweetHandler {
	return CreateTweetHandler{TweetService: tweetService}
}

func (h CreateTweetHandler) Handle(ctx context.Context, cmd CreateTweet) error {
	h.TweetService.CreateTweet(ctx, cmd.Tweet)
	return nil
}
