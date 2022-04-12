package command

import (
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
