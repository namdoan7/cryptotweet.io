package app

import (
	"github.com/levinhne/cryptotweet.io/internal/tag/app/command"
)

type Commands struct {
	FindOrCreate command.FindOrCreateTagHandler
}

type Application struct {
	Commands Commands
}
