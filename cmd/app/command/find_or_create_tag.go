package command

import (
	"context"
)

type FinOrCreateTag struct {
	Name string
}

type FindOrCreateTagHandler struct {
	TagService TagService
}

func NewFindOrCreateTagHandler(tagService TagService) FindOrCreateTagHandler {
	return FindOrCreateTagHandler{TagService: tagService}
}

func (h FindOrCreateTagHandler) Handle(ctx context.Context, cmd FinOrCreateTag) error {
	h.TagService.FindOrCreate(ctx, cmd.Name)
	return nil
}
