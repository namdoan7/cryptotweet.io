package adapters

import (
	"context"
	"log"

	tweetpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tweet"
	tweet "github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TweetGrpc struct {
	client tweetpb.TweetServiceClient
}

func NewTweetGrpc(client tweetpb.TweetServiceClient) TweetGrpc {
	return TweetGrpc{client: client}
}

func (s TweetGrpc) CreateTweet(ctx context.Context, tweet *tweet.Tweet) error {
	var photos []*tweetpb.Photo
	for _, photo := range tweet.Photos {
		photos = append(photos, &tweetpb.Photo{
			Width:       photo.Width,
			Height:      photo.Height,
			Url:         photo.Url,
			ExpandedUrl: photo.ExpandedUrl,
		})
	}

	var entities tweetpb.Entities
	// hashtag
	for _, hashtag := range tweet.Entities.Hashtags {
		entities.Hashtags = append(entities.Hashtags, &tweetpb.Entity{
			Text:        hashtag.Text,
			Url:         hashtag.Url,
			Indices:     hashtag.Indices,
			DisplayUrl:  hashtag.DisplayUrl,
			ExpandedUrl: hashtag.DisplayUrl,
		})
	}

	_, err := s.client.Create(ctx, &tweetpb.CreateTweetRequest{
		TweetId:          tweet.TweetId,
		TwitterProfileId: tweet.TwitterProfileId,
		Text:             tweet.Text,
		TranslateText: &tweetpb.TranslateText{
			Vietnamese: tweet.TranslateText.Vietnamese,
			Russian:    tweet.TranslateText.Russian,
		},
		FavoriteCount:        tweet.FavoriteCount,
		ConversationCount:    tweet.ConversationCount,
		Lang:                 tweet.Lang,
		TweetedAt:            &timestamppb.Timestamp{Seconds: tweet.TweetedAt.Unix()},
		InReplyToScreenName:  tweet.InReplyToScreenName,
		InReplyToStatusIdStr: tweet.InReplyToStatusIdStr,
		InReplyToUserIdStr:   tweet.InReplyToUserIdStr,
		Photos:               photos,
		Entities:             &entities,
	})
	log.Println(err)
	return err
}
