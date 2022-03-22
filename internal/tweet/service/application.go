package service

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tweet/adapters"
	"github.com/levinhne/cryptotweet.io/internal/tweet/app"
	"github.com/levinhne/cryptotweet.io/internal/tweet/app/command"
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
	hourRepository := adapters.NewMongoTweetRepository(client.Database("tweets"))

	return app.Application{
		Commands: app.Commands{
			CreateTweet: *command.NewCreateTweetHandler(hourRepository),
		},
	}
}
