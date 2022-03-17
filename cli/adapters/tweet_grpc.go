package adapters

import (
	"context"

	tweetpb "github.com/cryptotweet.io/internal/common/genproto/tweet"
	tweet "github.com/cryptotweet.io/internal/tweet/domain/tweet"
)

type TweetGrpc struct {
	client tweetpb.TweetServiceClient
}

func NewTweetGrpc(client tweetpb.TweetServiceClient) TweetGrpc {
	return TweetGrpc{client: client}
}

func (s TweetGrpc) CreateTweet(ctx context.Context, tweet tweet.Tweet) error {
	_, err := s.client.Create(ctx, &tweetpb.CreateTweetRequest{})
	return err
}
