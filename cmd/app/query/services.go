package query

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
)

type TagService interface {
	GetTag(ctx context.Context, name string) (*tag.Tag, error)
}
