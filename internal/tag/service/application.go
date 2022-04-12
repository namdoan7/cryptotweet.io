package service

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tag/adapters"
	"github.com/levinhne/cryptotweet.io/internal/tag/app"
	"github.com/levinhne/cryptotweet.io/internal/tag/app/command"
	"github.com/levinhne/cryptotweet.io/internal/tag/app/query"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewApplication() app.Application {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/?authSource=admin"),
	)
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}
	repository := adapters.NewMongoTagRepository(client.Database("tweets"))

	return app.Application{
		Commands: app.Commands{
			FindOrCreate: *command.NewFindOrCreateTagHandler(repository),
			CreateTag:    *command.NewCreateTagHandler(repository),
		},
		Queries: app.Queries{
			GetTag: *query.NewGetTagHandler(repository),
		},
	}
}
