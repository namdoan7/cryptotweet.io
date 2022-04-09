package adapters

import (
	"context"
	"log"
	"time"

	"github.com/chidiwilliams/flatbson"
	"github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTweetRepository struct {
	MongoDB *mongo.Database
}

func NewMongoTweetRepository(mongodb *mongo.Database) *MongoTweetRepository {
	return &MongoTweetRepository{MongoDB: mongodb}
}

func (m MongoTweetRepository) Find() ([]tweet.Tweet, error) {
	cursor, err := m.MongoDB.Collection("tweets").Find(context.Background(), bson.D{})
	log.Println("vinh", err)
	var tweets = make([]tweet.Tweet, 0)
	for cursor.Next(context.Background()) {
		var tweet tweet.Tweet
		cursor.Decode(&tweet)
		tweets = append(tweets, tweet)
	}
	cursor.Decode(&tweets)
	return tweets, err
}

func (m MongoTweetRepository) Create(document tweet.Tweet) error {
	_, err := m.MongoDB.Collection("tweets").InsertOne(context.Background(), document)
	return err
}

func (m MongoTweetRepository) Update(document tweet.Tweet) error {
	document.UpdatedAt = time.Now()
	update, err := flatbson.Flatten(document)
	if err != nil {
		return err
	}
	_, err = m.MongoDB.Collection("tweets").UpdateOne(
		context.Background(),
		bson.M{"tweet_id": document.TweetId},
		bson.M{"$set": update},
	)
	return err
}
