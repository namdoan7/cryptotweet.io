package main

import (
	tweetpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tweet"
	"github.com/levinhne/cryptotweet.io/internal/common/server"
	"github.com/levinhne/cryptotweet.io/internal/tweet/ports"
	"github.com/levinhne/cryptotweet.io/internal/tweet/service"
	"google.golang.org/grpc"
)

func main() {
	application := service.NewApplication()

	server.RunGRPCServer(func(server *grpc.Server) {
		srv := ports.NewGrpcServer(application)
		tweetpb.RegisterTweetServiceServer(server, srv)
	})
}
