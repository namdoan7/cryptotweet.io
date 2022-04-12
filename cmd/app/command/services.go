package command

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/profile/domain/profile"
	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
	"github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
)

type TweetService interface {
	CreateTweet(ctx context.Context, tweet *tweet.Tweet) error
	UpdateTweet(ctx context.Context, tweet *tweet.Tweet) error
}

type ProfileService interface {
	CreateProfile(ctx context.Context, profile *profile.Profile) error
}

type TagService interface {
	GetTag(ctx context.Context, name string) (*tag.Tag, error)
	CreateTag(ctx context.Context, tag tag.Tag) error
}
