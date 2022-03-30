package main

import (
	profilepb "github.com/levinhne/cryptotweet.io/internal/common/genproto/profile"
	"github.com/levinhne/cryptotweet.io/internal/common/server"
	"github.com/levinhne/cryptotweet.io/internal/profile/ports"
	"github.com/levinhne/cryptotweet.io/internal/profile/service"
	"google.golang.org/grpc"
)

func main() {
	application := service.NewApplication()

	server.RunGRPCServer(func(server *grpc.Server) {
		srv := ports.NewGrpcServer(application)
		profilepb.RegisterProfileServiceServer(server, srv)
	})
}
