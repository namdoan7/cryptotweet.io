package app

import "github.com/cryptotweet.io/internal/tweet/app/command"

type Commands struct {
	CreateTweet command.CreateTweetHandler
}

type Application struct {
	Commands Commands
}
