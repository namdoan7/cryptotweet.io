package app

import "github.com/levinhne/cryptotweet.io/cmd/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	CreateTweet   command.CreateTweetHandler
	UpdateTweet   command.UpdateTweetHandler
	CreateProfile command.CreateProfileHandler
	CreateTag     command.CreateTagHandler
	// FinOrCreateTag command.FindOrCreateTagHandler
}
