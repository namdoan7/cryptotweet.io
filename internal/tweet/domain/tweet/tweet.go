package tweet

import "time"

type Entity struct {
	Indices     []int32 `bson:"indices,omitempty"`
	Text        string  `bson:"text,omitempty"`
	Url         string  `bson:"url,omitempty"`
	DisplayUrl  string  `bson:"display_url,omitempty"`
	ExpandedUrl string  `bson:"expanded_url,omitempty"`
}

type Entities struct {
	Hashtags     Entity `bson:"hashtags,omitempty"`
	Symbols      Entity `bson:"symbols,omitempty"`
	Media        Entity `bson:"media,omitempty"`
	Urls         Entity `bson:"urls,omitempty"`
	UserMentions Entity `bson:"user_mentions,omitempty"`
}

type Photo struct {
	Width       int32  `bson:"width,omitempty"`
	Height      int32  `bson:"height,omitempty"`
	Url         string `bson:"url,omitempty"`
	ExpandedUrl string `bson:"expanded_url,omitempty"`
}

type Photos struct {
	Photos []*Photo `bson:"photos,omitempty"`
}

type Text struct {
	En string `bson:"en,omitempty"`
	Vi string `bson:"vi,omitempty"`
}

type Tweet struct {
	Id                string `bson:"_id,omitempty"`
	TweetId           string `bson:"tweet_id,omitempty"`
	TwitterProfileId  string `bson:"twitter_profile_id,omitempty"`
	Text              *Text  `bson:"text,omitempty"`
	FavoriteCount     int32  `bson:"favorite_count,omitempty"`
	ConversationCount int32  `bson:"conversation_count,omitempty"`
	Lang              string `bson:"lang,omitempty"`
	// Parent               Tweet     `bson:"parent,omitempty"`
	Photos               Photos    `bson:"photos,omitempty"`
	Entities             Entities  `bson:"entities,omitempty"`
	InReplyToScreenName  string    `bson:"in_reply_to_screen_name,omitempty"`
	InReplyToStatusIdStr string    `bson:"in_reply_to_status_id_str,omitempty"`
	InReplyToUserIdStr   string    `bson:"in_reply_to_user_id_str,omitempty"`
	TweetedAt            time.Time `bson:"tweeted_at,omitempty"`
	PublishedAt          time.Time `bson:"published_at,omitempty"`
	UpdatedAt            time.Time `bson:"updated_at,omitempty"`
}
