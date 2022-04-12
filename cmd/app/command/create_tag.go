package command

import (
	"context"

	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
)

type CreateTag struct {
	Tag tag.Tag
}

type CreateTagHandler struct {
	TagService TagService
}

func NewCreateTagHandler(tagService TagService) CreateTagHandler {
	return CreateTagHandler{TagService: tagService}
}

func (h CreateTagHandler) Handle(ctx context.Context, cmd CreateTag) error {
	return h.TagService.CreateTag(ctx, cmd.Tag)
}
