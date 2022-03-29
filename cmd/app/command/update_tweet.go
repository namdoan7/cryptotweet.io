package command

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
)

type UpdateTweet struct {
	Tweet *tweet.Tweet
}

type UpdateTweetHandler struct {
	TweetService TweetService
}

func NewUpdateTweetHandler(tweetService TweetService) UpdateTweetHandler {
	return UpdateTweetHandler{TweetService: tweetService}
}

func (h UpdateTweetHandler) Handle(ctx context.Context, cmd CreateTweet) error {
	h.TweetService.UpdateTweet(ctx, cmd.Tweet)
	return nil
}
