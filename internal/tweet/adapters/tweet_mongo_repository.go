package adapters

import (
	"context"
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

func (m MongoTweetRepository) Find(filter bson.M) ([]*tweet.Tweet, error) {
	tweets := make([]*tweet.Tweet, 0)
	cursor, err := m.MongoDB.Collection("tweets").Find(context.Background(), filter)
	if err != nil {
		return tweets, err
	}
	err = cursor.All(context.Background(), &tweets)
	return tweets, err
}

func (m MongoTweetRepository) GetTweet(tweetId string) (*tweet.Tweet, error) {
	result := m.MongoDB.Collection("tweets").FindOne(context.Background(), bson.M{"tweet_id": tweetId})
	var tweet tweet.Tweet
	err := result.Decode(&tweet)
	return &tweet, err
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
