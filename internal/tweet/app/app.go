package app

import (
	"github.com/levinhne/cryptotweet.io/internal/tweet/app/command"
	"github.com/levinhne/cryptotweet.io/internal/tweet/app/query"
)

type Commands struct {
	CreateTweet command.CreateTweetHandler
	UpdateTweet command.UpdateTweetHandler
}

type Queries struct {
	ListTweets query.ListTweetsHandler
	GetTweet   query.GetTweetHandler
}

type Application struct {
	Commands Commands
	Queries  Queries
}
