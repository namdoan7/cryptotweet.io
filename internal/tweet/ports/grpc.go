package ports

import (
	"context"
	"encoding/json"
	"log"

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
	tweet := trans[tweet.Tweet](in)
	g.app.Commands.CreateTweet.Handle(tweet)
	return &tweetpb.CreateTweetResponse{}, nil
}

func (g GrpcServer) Update(ctx context.Context, in *tweetpb.UpdateTweetRequest) (*tweetpb.UpdateTweetResponse, error) {
	tweet := trans[tweet.Tweet](in)
	log.Println(in.String())
	g.app.Commands.UpdateTweet.Handle(tweet)
	return &tweetpb.UpdateTweetResponse{}, nil
}
