package models

import (
	"time"
)

type Entity struct {
	Indices     []int32 `bson:"indices,omitempty" json:"indices,omitempty"`
	Text        string  `bson:"text,omitempty" json:"text,omitempty"`
	Url         string  `bson:"url,omitempty" json:"url,omitempty"`
	DisplayUrl  string  `bson:"display_url,omitempty" json:"display_url,omitempty"`
	ExpandedUrl string  `bson:"expanded_url,omitempty" json:"expanded_url,omitempty"`
}

type Profile struct {
	Id               string    `bson:"_id,omitempty"`
	ProfileTwitterId string                 `bson:"profile_twitter_id,omitempty"`
	Name             string                 `bson:"name,omitempty"`
	ScreenName       string                 `bson:"screen_name,omitempty"`
	FavouritesCount  int32                  `bson:"favourites_count,omitempty"`
	FollowersCount   int32                  `bson:"followers_count,omitempty"`
	Verified         bool                   `bson:"verified,omitempty"`
	Description      string                 `bson:"description,omitempty"`
	ProfileImageUrl  string                 `bson:"profile_image_url,omitempty"`
	Entities         map[string]interface{} `bson:"entities,omitempty"`
	PinnedTweetIds []string `bson:"pinned_tweet_ids,omitempty"`
	CreatedAt        time.Time              `bson:"created_at,omitempty"`
	UpdatedAt        time.Time              `bson:"updated_at,omitempty"`
}