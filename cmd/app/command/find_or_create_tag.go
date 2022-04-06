package command

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
)

type FindOrCreateTag struct {
	Tag tag.Tag
}

type FindOrCreateTagHandler struct {
	TagService TagService
}

func NewFindOrCreateTagHandler(tagService TagService) FindOrCreateTagHandler {
	return FindOrCreateTagHandler{TagService: tagService}
}

func (h FindOrCreateTagHandler) Handle(ctx context.Context, cmd FindOrCreateTag) (*tag.Tag, error) {
	tag, err := h.TagService.FindOrCreate(ctx, cmd.Tag)
	return tag, err
}
