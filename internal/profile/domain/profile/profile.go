package profile

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

type EntityDescription struct {
	Urls []Entity `bson:"urls,omitempty" json:"urls,omitempty"`
}

type EntiryUrl struct {
	Urls []Entity `bson:"urls,omitempty" json:"urls,omitempty"`
}

type Entities struct {
	Description EntityDescription `bson:"description,omitempty" json:"description,omitempty"`
	Url         EntiryUrl         `bson:"url,omitempty" json:"url,omitempty"`
}

type Profile struct {
	Id               string    `bson:"_id,omitempty"`
	ProfileTwitterId string    `bson:"profile_twitter_id,omitempty" json:"profile_twitter_id,omitempty"`
	Name             string    `bson:"name,omitempty" json:"name,omitempty"`
	ScreenName       string    `bson:"screen_name,omitempty" json:"screen_name,omitempty"`
	FavouritesCount  int32     `bson:"favourites_count,omitempty" json:"favourites_count,omitempty"`
	FollowersCount   int32     `bson:"followers_count,omitempty" json:"followers_count,omitempty"`
	FriendsCount     int32     `bson:"friends_count,omitempty" json:"friends_count,omitempty"`
	Verified         bool      `bson:"verified,omitempty" json:"verified,omitempty"`
	Description      string    `bson:"description,omitempty" json:"description,omitempty"`
	Entities         Entities  `bson:"entities,omitempty" json:"entities,omitempty"`
	ProfileImageUrl  string    `bson:"profile_image_url,omitempty" json:"profile_image_url,omitempty"`
	ProfileBannerUrl string    `bson:"profile_banner_url,omitempty" json:"profile_banner_url,omitempty"`
	PinnedTweetIds   []string  `bson:"pinned_tweet_ids,omitempty" json:"pinned_tweet_ids,omitempty"`
	CreatedAt        time.Time `bson:"created_at,omitempty"`
	UpdatedAt        time.Time `bson:"updated_at,omitempty"`
}
