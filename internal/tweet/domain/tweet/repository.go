package tweet

import "go.mongodb.org/mongo-driver/bson"

type Repository interface {
	Find(filter bson.M) ([]*Tweet, error)
	GetTweet(tweetId string) (*Tweet, error)
	Create(tweet Tweet) error
	Update(tweet Tweet) error
}
