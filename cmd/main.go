package main

import (
	"context"
	"log"

	"github.com/levinhne/cryptotweet.io/cmd/cmd"
	"github.com/levinhne/cryptotweet.io/cmd/service"
)

func main() {
	ctx := context.Background()
	app, cleanup := service.NewApplication(ctx)
	defer cleanup()
	err := cmd.Execute(app)
	if err != nil {
		log.Fatal(err)
	}
}
