package adapters

import (
	"context"
	"log"
	"time"

	"github.com/chidiwilliams/flatbson"
	"github.com/levinhne/cryptotweet.io/internal/profile/domain/profile"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProfileRepository struct {
	MongoDB *mongo.Database
}

func NewMongoProfileRepository(mongodb *mongo.Database) *MongoProfileRepository {
	return &MongoProfileRepository{MongoDB: mongodb}
}

func (m MongoProfileRepository) Find() ([]profile.Profile, error) {
	cursor, err := m.MongoDB.Collection("profiles").Find(context.Background(), bson.M{})
	if err != nil {
		return make([]profile.Profile, 0), err
	}
	var profiles []profile.Profile
	if err = cursor.All(context.Background(), &profiles); err != nil {
		log.Fatal(err)
	}
	log.Println(profiles)
	return profiles, err
}

func (m MongoProfileRepository) Create(document profile.Profile) error {
	_, err := m.MongoDB.Collection("profiles").InsertOne(context.Background(), document)
	return err
}

func (m MongoProfileRepository) Update(document profile.Profile) error {
	document.UpdatedAt = time.Now()
	update, err := flatbson.Flatten(document)
	if err != nil {
		return err
	}
	_, err = m.MongoDB.Collection("profiles").UpdateOne(
		context.Background(),
		bson.M{"profile_id": document.ProfileTwitterId},
		bson.M{"$set": update},
	)
	return err
}
