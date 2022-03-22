package app

import "github.com/levinhne/cryptotweet.io/cmd/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	CreateTweet command.CreateTweetHandler
}
