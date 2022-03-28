package adapters

import (
	"context"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
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
	ee, _ := json.Marshal(tweet)
	var createTweetRequest tweetpb.CreateTweetRequest
	jsonpb.Unmarshal(ee, &createTweetRequest)

	// var photos []*tweetpb.Photo
	// for _, photo := range tweet.Photos {
	// 	photos = append(photos, &tweetpb.Photo{
	// 		Width:       photo.Width,
	// 		Height:      photo.Height,
	// 		Url:         photo.Url,
	// 		ExpandedUrl: photo.ExpandedUrl,
	// 	})
	// }

	// var entities tweetpb.Entities
	// // hashtag
	// for _, hashtag := range tweet.Entities.Hashtags {
	// 	entities.Hashtags = append(entities.Hashtags, &tweetpb.Entity{
	// 		Text:        hashtag.Text,
	// 		Url:         hashtag.Url,
	// 		Indices:     hashtag.Indices,
	// 		DisplayUrl:  hashtag.DisplayUrl,
	// 		ExpandedUrl: hashtag.DisplayUrl,
	// 	})
	// }

	_, err := s.client.Create(ctx, &createTweetRequest)
	return err
}
