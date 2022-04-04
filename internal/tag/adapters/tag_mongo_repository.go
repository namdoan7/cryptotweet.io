package adapters

import (
	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTagRepository struct {
	MongoDB *mongo.Database
}

func NewMongoTagRepository(mongodb *mongo.Database) *MongoTagRepository {
	return &MongoTagRepository{MongoDB: mongodb}
}

func (m MongoTagRepository) FindOrCreate(name string) (tag.Tag, error) {
	return tag.Tag{}, nil
}
