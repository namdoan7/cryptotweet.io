package command

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
)

type TweetService interface {
	CreateTweet(ctx context.Context, tweet *tweet.Tweet) error
}
