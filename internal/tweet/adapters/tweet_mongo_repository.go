package adapters

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTweetRepository struct {
	MongoDB *mongo.Database
}

func NewMongoTweetRepository(mongodb *mongo.Database) *MongoTweetRepository {
	return &MongoTweetRepository{MongoDB: mongodb}
}

func (m MongoTweetRepository) Create(document tweet.Tweet) error {
	_, err := m.MongoDB.Collection("tweets").InsertOne(context.Background(), document)
	return err
}
