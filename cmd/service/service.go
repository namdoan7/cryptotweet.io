package service

import (
	"context"

	"github.com/levinhne/cryptotweet.io/cmd/adapters"
	"github.com/levinhne/cryptotweet.io/cmd/app"
	"github.com/levinhne/cryptotweet.io/cmd/app/command"
	"github.com/levinhne/cryptotweet.io/internal/common/client"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	// tweet client
	tweetClient, closeTweetClient, _ := client.NewTweetClient()
	tweetGrpc := adapters.NewTweetGrpc(tweetClient)

	// profile client
	profileClient, closeProfileClient, _ := client.NewProfileClient()
	profileGrpc := adapters.NewProfileGrpc(profileClient)

	// taf client
	tagClient, closeTagClient, _ := client.NewTagClient()
	tagGrpc := adapters.NewTagGrpc(tagClient)

	return newApplication(ctx, tweetGrpc, profileGrpc, tagGrpc),
		func() {
			closeTweetClient()
			closeProfileClient()
			closeTagClient()
		}
}

func newApplication(
	ctx context.Context,
	tweetGrpc command.TweetService,
	profileGrpc command.ProfileService,
	tagGrpc command.TagService,
) app.Application {
	return app.Application{
		Commands: app.Commands{
			CreateTweet:    command.NewCreateTweetHandler(tweetGrpc),
			UpdateTweet:    command.NewUpdateTweetHandler(tweetGrpc),
			CreateProfile:  command.NewCreateProfileHandler(profileGrpc),
			FinOrCreateTag: command.NewFindOrCreateTagHandler(tagGrpc),
		},
	}
}
