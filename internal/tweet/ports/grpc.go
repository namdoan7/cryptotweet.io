package ports

import (
	"context"

	tweetpb "github.com/cryptotweet.io/internal/common/genproto/tweet"
	"github.com/cryptotweet.io/internal/tweet/app"
	"github.com/cryptotweet.io/internal/tweet/domain/tweet"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) Create(ctx context.Context, in *tweetpb.CreateTweetRequest) (*tweetpb.CreateTweetResponse, error) {
	g.app.Commands.CreateTweet.Handle(tweet.Tweet{})

	return &tweetpb.CreateTweetResponse{}, nil
}
