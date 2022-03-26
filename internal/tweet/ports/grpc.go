package ports

import (
	"context"

	"github.com/k0kubun/pp/v3"
	tweetpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tweet"
	"github.com/levinhne/cryptotweet.io/internal/tweet/app"
	"github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
)

type GrpcServer struct {
	app app.Application
	tweetpb.UnimplementedTweetServiceServer
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) Create(ctx context.Context, in *tweetpb.CreateTweetRequest) (*tweetpb.CreateTweetResponse, error) {
	tweet := tweet.Tweet{
		TweetId:          in.TweetId,
		Text:             in.Text,
		TwitterProfileId: in.TwitterProfileId,
		TranslateText: tweet.TranslateText{
			Vietnamese: in.TranslateText.Vietnamese,
			Russian:    in.TranslateText.Russian,
		},
		FavoriteCount:       in.FavoriteCount,
		ConversationCount:   in.ConversationCount,
		Lang:                in.Lang,
		InReplyToScreenName: in.InReplyToScreenName,
	}
	pp.Println(tweet)
	g.app.Commands.CreateTweet.Handle(tweet)
	return &tweetpb.CreateTweetResponse{}, nil
}
