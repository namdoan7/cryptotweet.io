package adapters

import (
	"context"

	tweet "github.com/cryptotweet.io/internal/tweet/domain/tweet"
)

type TweetService interface {
	Create(ctx context.Context, tweet tweet.Tweet) error
}
