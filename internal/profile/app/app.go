package app

import (
	"github.com/levinhne/cryptotweet.io/internal/profile/app/command"
	"github.com/levinhne/cryptotweet.io/internal/profile/app/query"
)

type Commands struct {
	CreateProfile command.CreateProfileHandler
}

type Queries struct {
	FindProfile query.FindProfileHandler
}

type Application struct {
	Commands Commands
	Queries  Queries
}
