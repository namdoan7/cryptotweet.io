package app

import (
	"github.com/levinhne/cryptotweet.io/cmd/app/command"
	"github.com/levinhne/cryptotweet.io/cmd/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateTweet   command.CreateTweetHandler
	UpdateTweet   command.UpdateTweetHandler
	CreateProfile command.CreateProfileHandler
	CreateTag     command.CreateTagHandler
	// FinOrCreateTag command.FindOrCreateTagHandler
}

type Queries struct {
	GetTag query.GetTagHandler
}
