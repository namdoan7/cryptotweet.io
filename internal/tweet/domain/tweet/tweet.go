package tweet

import "time"

type Entity struct {
	Indices     []uint16 `bson:"indices,omitempty" json:"indices,omitempty"`
	Text        string   `bson:"text,omitempty" json:"text,omitempty"`
	Url         string   `bson:"url,omitempty" json:"url,omitempty"`
	DisplayUrl  string   `bson:"display_url,omitempty" json:"display_url,omitempty"`
	ExpandedUrl string   `bson:"expanded_url,omitempty" json:"expanded_url,omitempty"`
}

type Entities struct {
	Hashtags     []Entity `bson:"hashtags,omitempty" json:"hashtags,omitempty"`
	Symbols      []Entity `bson:"symbols,omitempty" json:"symbols,omitempty"`
	Media        []Entity `bson:"media,omitempty" json:"media,omitempty"`
	Urls         []Entity `bson:"urls,omitempty" json:"urls,omitempty"`
	UserMentions []Entity `bson:"user_mentions,omitempty" json:"user_mentions,omitempty"`
}

type Photo struct {
	Width       uint16 `bson:"width,omitempty" json:"width,omitempty"`
	Height      uint16 `bson:"height,omitempty" json:"height,omitempty"`
	Url         string `bson:"url,omitempty" json:"url,omitempty"`
	ExpandedUrl string `bson:"expanded_url,omitempty" json:"expanded_url,omitempty"`
}

type TranslateText struct {
	Vietnamese string `bson:"vi,omitempty"`
	Russian    string `bson:"ru,omitempty"`
}

type Tweet struct {
	Id                   string        `bson:"_id,omitempty"`
	TweetId              string        `bson:"tweet_id,omitempty" json:"tweet_id,omitempty"`
	TwitterProfileId     string        `bson:"twitter_profile_id,omitempty" json:"twitter_profile_id,omitempty"`
	Text                 string        `bson:"text,omitempty" json:"text,omitempty"`
	TranslateText        TranslateText `bson:"translate_text,omitempty" json:"translate_text,omitempty"`
	FavoriteCount        int32         `bson:"favorite_count,omitempty" json:"favorite_count,omitempty"`
	ConversationCount    int32         `bson:"conversation_count,omitempty" json:"conversation_count,omitempty"`
	Lang                 string        `bson:"lang,omitempty" json:"lang,omitempty"`
	Hashtags             []int32       `bson:"hashtags,omitempty" json:"hashtags,omitempty"`
	Photos               []Photo       `bson:"photos,omitempty" json:"photos,omitempty"`
	Entities             Entities      `bson:"entities,omitempty" json:"entities,omitempty"`
	InReplyToScreenName  string        `bson:"in_reply_to_screen_name,omitempty" json:"in_reply_to_screen_name,omitempty"`
	InReplyToStatusIdStr string        `bson:"in_reply_to_status_id_str,omitempty" json:"in_reply_to_status_id_str,omitempty"`
	InReplyToUserIdStr   string        `bson:"in_reply_to_user_id_str,omitempty" json:"in_reply_to_user_id_str,omitempty"`
	TweetedAt            time.Time     `bson:"tweeted_at,omitempty" json:"tweeted_at,omitempty"`
	PublishedAt          time.Time     `bson:"published_at,omitempty" json:"published_at,omitempty"`
	UpdatedAt            time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	CreatedAt            time.Time     `bson:"create_at,omitempty" json:"create_at,omitempty"`
	// Parent               Tweet     `bson:"parent,omitempty" json:"parent,omitempty"`
}
