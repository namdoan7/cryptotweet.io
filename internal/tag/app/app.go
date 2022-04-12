package app

import (
	"github.com/levinhne/cryptotweet.io/internal/tag/app/command"
	"github.com/levinhne/cryptotweet.io/internal/tag/app/query"
)

type Commands struct {
	FindOrCreate command.FindOrCreateTagHandler
	CreateTag    command.CreateTagHandler
}

type Queries struct {
	GetTag query.GetTagHandler
}

type Application struct {
	Commands Commands
	Queries  Queries
}
