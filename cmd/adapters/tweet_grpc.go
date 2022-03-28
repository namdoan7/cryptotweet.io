package adapters

import (
	"context"
	"encoding/json"

	tweetpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tweet"
	tweet "github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
)

type TweetGrpc struct {
	client tweetpb.TweetServiceClient
}

func NewTweetGrpc(client tweetpb.TweetServiceClient) TweetGrpc {
	return TweetGrpc{client: client}
}

func (s TweetGrpc) CreateTweet(ctx context.Context, tweet *tweet.Tweet) error {
	ee, err := json.Marshal(tweet)
	var createTweetRequest tweetpb.CreateTweetRequest
	err = json.Unmarshal(ee, &createTweetRequest)
	_, err = s.client.Create(ctx, &createTweetRequest)
	return err
}
