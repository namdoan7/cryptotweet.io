package tweet

import (
	"time"

	tweetpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tweet"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Entity struct {
	Indices     []uint32 `bson:"indices,omitempty" json:"indices,omitempty"`
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
	Width       uint32 `bson:"width,omitempty" json:"width,omitempty"`
	Height      uint32 `bson:"height,omitempty" json:"height,omitempty"`
	Url         string `bson:"url,omitempty" json:"url,omitempty"`
	ExpandedUrl string `bson:"expanded_url,omitempty" json:"expanded_url,omitempty"`
}

type TranslateText struct {
	Vietnamese string `bson:"vi,omitempty" json:"vi,omitempty"`
	Russian    string `bson:"ru,omitempty" json:"ru,omitempty"`
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

func (t *Tweet) ToProtoMessage() *tweetpb.Tweet {
	var photos = make([]*tweetpb.Photo, 0)
	for _, photo := range t.Photos {
		photos = append(photos, &tweetpb.Photo{
			Width:       photo.Width,
			Height:      photo.Height,
			Url:         photo.Url,
			ExpandedUrl: photo.ExpandedUrl,
		})
	}
	var hashtags = make([]*tweetpb.Entity, 0)
	for _, entity := range t.Entities.Hashtags {
		hashtags = append(hashtags, &tweetpb.Entity{
			Indices:     entity.Indices,
			Text:        entity.Text,
			Url:         entity.Url,
			DisplayUrl:  entity.DisplayUrl,
			ExpandedUrl: entity.ExpandedUrl,
		})
	}

	var symbols = make([]*tweetpb.Entity, 0)
	for _, entity := range t.Entities.Symbols {
		symbols = append(symbols, &tweetpb.Entity{
			Indices:     entity.Indices,
			Text:        entity.Text,
			Url:         entity.Url,
			DisplayUrl:  entity.DisplayUrl,
			ExpandedUrl: entity.ExpandedUrl,
		})
	}
	var media = make([]*tweetpb.Entity, 0)
	for _, entity := range t.Entities.Media {
		media = append(media, &tweetpb.Entity{
			Indices:     entity.Indices,
			Text:        entity.Text,
			Url:         entity.Url,
			DisplayUrl:  entity.DisplayUrl,
			ExpandedUrl: entity.ExpandedUrl,
		})
	}
	var urls = make([]*tweetpb.Entity, 0)
	for _, entity := range t.Entities.Urls {
		urls = append(urls, &tweetpb.Entity{
			Indices:     entity.Indices,
			Text:        entity.Text,
			Url:         entity.Url,
			DisplayUrl:  entity.DisplayUrl,
			ExpandedUrl: entity.ExpandedUrl,
		})
	}
	var userMentions = make([]*tweetpb.Entity, 0)
	for _, entity := range t.Entities.UserMentions {
		userMentions = append(userMentions, &tweetpb.Entity{
			Indices:     entity.Indices,
			Text:        entity.Text,
			Url:         entity.Url,
			DisplayUrl:  entity.DisplayUrl,
			ExpandedUrl: entity.ExpandedUrl,
		})
	}
	return &tweetpb.Tweet{
		Id:               t.Id,
		TweetId:          t.TweetId,
		TwitterProfileId: t.TwitterProfileId,
		Text:             t.Text,
		TranslateText: &tweetpb.TranslateText{
			Vietnamese: t.TranslateText.Vietnamese,
			Russian:    t.TranslateText.Russian,
		},
		Lang:   t.Lang,
		Photos: photos,
		Entities: &tweetpb.Entities{
			Hashtags:     hashtags,
			Symbols:      symbols,
			Media:        media,
			Urls:         urls,
			UserMentions: userMentions,
		},
		FavoriteCount:        t.FavoriteCount,
		ConversationCount:    t.ConversationCount,
		InReplyToScreenName:  t.InReplyToScreenName,
		InReplyToStatusIdStr: t.InReplyToStatusIdStr,
		InReplyToUserIdStr:   t.InReplyToUserIdStr,
		TweetedAt:            &timestamppb.Timestamp{Seconds: t.TweetedAt.Unix()},
	}
}
