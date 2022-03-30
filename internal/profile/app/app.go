package app

import "github.com/levinhne/cryptotweet.io/internal/profile/app/command"

type Commands struct {
	CreateProfile command.CreateProfileHandler
}

type Application struct {
	Commands Commands
}
