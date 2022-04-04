package main

import (
	tagpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tag"
	"github.com/levinhne/cryptotweet.io/internal/common/server"
	"github.com/levinhne/cryptotweet.io/internal/tag/ports"
	"github.com/levinhne/cryptotweet.io/internal/tag/service"
	"google.golang.org/grpc"
)

func main() {
	application := service.NewApplication()

	server.RunGRPCServer(func(server *grpc.Server) {
		srv := ports.NewGrpcServer(application)
		tagpb.RegisterTagServiceServer(server, srv)
	})
}
