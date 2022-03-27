package ports

import (
	"context"
	"encoding/json"

	tweetpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tweet"
	"github.com/levinhne/cryptotweet.io/internal/tweet/app"
	"github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func trans[R any](s protoreflect.ProtoMessage) R {
	m := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	var r R
	ee, _ := m.Marshal(s)
	json.Unmarshal(ee, &r)
	return r
}

type GrpcServer struct {
	app app.Application
	tweetpb.UnimplementedTweetServiceServer
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) Create(ctx context.Context, in *tweetpb.CreateTweetRequest) (*tweetpb.CreateTweetResponse, error) {
	t2 := trans[tweet.Tweet](in)
	// for _, photo := range in.Photos {
	// 	photos = append(photos, tweet.Photo{
	// 		Width:       photo.Width,
	// 		Height:      photo.Height,
	// 		Url:         photo.Url,
	// 		ExpandedUrl: photo.ExpandedUrl,
	// 	})
	// }

	// var entities tweet.Entities
	// // hashtag
	// for _, hashtag := range in.Entities.Hashtags {
	// 	entities.Hashtags = append(entities.Hashtags, tweet.Entity{
	// 		Text:        hashtag.Text,
	// 		Url:         hashtag.Url,
	// 		Indices:     hashtag.Indices,
	// 		DisplayUrl:  hashtag.DisplayUrl,
	// 		ExpandedUrl: hashtag.DisplayUrl,
	// 	})
	// }
	// // media
	// for _, media := range in.Entities.Media {
	// 	entities.Media = append(entities.Media, tweet.Entity{
	// 		Text:        media.Text,
	// 		Url:         media.Url,
	// 		Indices:     media.Indices,
	// 		DisplayUrl:  media.DisplayUrl,
	// 		ExpandedUrl: media.DisplayUrl,
	// 	})
	// }
	// // symbol
	// for _, symbol := range in.Entities.Symbols {
	// 	entities.Symbols = append(entities.Symbols, tweet.Entity{
	// 		Text:        symbol.Text,
	// 		Url:         symbol.Url,
	// 		Indices:     symbol.Indices,
	// 		DisplayUrl:  symbol.DisplayUrl,
	// 		ExpandedUrl: symbol.ExpandedUrl,
	// 	})
	// }
	// // urls
	// for _, symbol := range in.Entities.Urls {
	// 	entities.Urls = append(entities.Urls, tweet.Entity{
	// 		Text:        symbol.Text,
	// 		Url:         symbol.Url,
	// 		Indices:     symbol.Indices,
	// 		DisplayUrl:  symbol.DisplayUrl,
	// 		ExpandedUrl: symbol.ExpandedUrl,
	// 	})
	// }
	// // user mentions
	// for _, mention := range in.Entities.UserMentions {
	// 	entities.UserMentions = append(entities.UserMentions, tweet.Entity{
	// 		Text:        mention.Text,
	// 		Url:         mention.Url,
	// 		Indices:     mention.Indices,
	// 		DisplayUrl:  mention.DisplayUrl,
	// 		ExpandedUrl: mention.DisplayUrl,
	// 	})
	// }
	// tweet := tweet.Tweet{
	// 	TweetId:          in.TweetId,
	// 	Text:             in.Text,
	// 	TwitterProfileId: in.TwitterProfileId,
	// 	TranslateText: tweet.TranslateText{
	// 		Vietnamese: in.TranslateText.Vietnamese,
	// 		Russian:    in.TranslateText.Russian,
	// 	},
	// 	FavoriteCount:        in.FavoriteCount,
	// 	ConversationCount:    in.ConversationCount,
	// 	Lang:                 in.Lang,
	// 	InReplyToScreenName:  in.InReplyToScreenName,
	// 	Photos:               photos,
	// 	Entities:             entities,
	// 	InReplyToStatusIdStr: in.InReplyToStatusIdStr,
	// 	TweetedAt:            in.TweetedAt.AsTime(),
	// 	InReplyToUserIdStr:   in.InReplyToUserIdStr,
	// }
	g.app.Commands.CreateTweet.Handle(t2)
	return &tweetpb.CreateTweetResponse{}, nil
}
