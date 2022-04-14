package query

import (
	"context"

	"github.com/levinhne/cryptotweet.io/cmd/adapters"
	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
)

type GetTagHandler struct {
	TagService adapters.TagService
}

func NewGetTagHandler(tagService adapters.TagService) GetTagHandler {
	return GetTagHandler{TagService: tagService}
}

func (h GetTagHandler) Handle(ctx context.Context, name string) (*tag.Tag, error) {
	return h.TagService.GetTag(ctx, name)
}
