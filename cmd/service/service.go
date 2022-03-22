package service

import (
	"context"

	"github.com/levinhne/cryptotweet.io/cmd/adapters"
	"github.com/levinhne/cryptotweet.io/cmd/app"
	"github.com/levinhne/cryptotweet.io/cmd/app/command"
	"github.com/levinhne/cryptotweet.io/internal/common/client"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	tweetClient, _, err := client.NewTweetClient()
	if err != nil {
		panic(err)
	}

	// usersClient, closeUsersClient, err := grpcClient.NewUsersClient()
	// if err != nil {
	// 	panic(err)
	// }
	tweetGrpc := adapters.NewTweetGrpc(tweetClient)
	// usersGrpc := adapters.NewUsersGrpc(usersClient)

	return newApplication(ctx, tweetGrpc),
		func() {
			// _ = closeTrainerClient()
			// _ = closeUsersClient()
		}
}

func newApplication(ctx context.Context, tweetGrpc command.TweetService) app.Application {
	return app.Application{
		Commands: app.Commands{
			CreateTweet: command.NewCreateTweetHandler(tweetGrpc),
		},
	}
}
