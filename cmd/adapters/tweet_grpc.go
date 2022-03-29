package adapters

import (
	"context"
	"encoding/json"
	"log"

	tweetpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tweet"
	tweet "github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
	"google.golang.org/protobuf/encoding/protojson"
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
	um := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	err = um.Unmarshal(ee, &createTweetRequest)
	_, err = s.client.Create(ctx, &createTweetRequest)
	return err
}

func (s TweetGrpc) UpdateTweet(ctx context.Context, tweet *tweet.Tweet) error {
	ee, err := json.Marshal(tweet)

	var updateTweetRequest tweetpb.UpdateTweetRequest
	um := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	err = um.Unmarshal(ee, &updateTweetRequest)
	log.Println(111111, updateTweetRequest.String())
	_, err = s.client.Update(ctx, &updateTweetRequest)
	return err
}
