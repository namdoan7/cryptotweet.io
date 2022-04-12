package command

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
)

type GetTag struct {
	Tag tag.Tag
}

type GetTagHandler struct {
	TagService TagService
}

func NewGetTagHandler(tagService TagService) *GetTagHandler {
	return &GetTagHandler{TagService: tagService}
}

func (h GetTagHandler) Handle(ctx context.Context, name string) (*tag.Tag, error) {
	return h.TagService.GetTag(ctx, name)
}
