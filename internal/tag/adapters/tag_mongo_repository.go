package adapters

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTagRepository struct {
	MongoDB *mongo.Database
}

func NewMongoTagRepository(mongodb *mongo.Database) *MongoTagRepository {
	return &MongoTagRepository{MongoDB: mongodb}
}

func (m MongoTagRepository) FindOrCreate(tag tag.Tag) (*tag.Tag, error) {
	filter := bson.M{"name": tag.Name}
	update := bson.M{
		"$set": bson.M{"name": tag.Name},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	result := m.MongoDB.Collection("tags").FindOneAndUpdate(context.Background(), filter, update, &opts)
	err := result.Decode(&tag)
	return &tag, err
}
