package command

import (
	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
)

type FindOrCreateTagHandler struct {
	TagRepository tag.Repository
}

func NewFindOrCreateTagHandler(tagRepository tag.Repository) FindOrCreateTagHandler {
	return FindOrCreateTagHandler{TagRepository: tagRepository}
}

func (h FindOrCreateTagHandler) Handle(tag tag.Tag) (*tag.Tag, error) {
	return h.TagRepository.FindOrCreate(tag)
}

type CreateTagHandler struct {
	TagRepository tag.Repository
}

func NewCreateTagHandler(tagRepository tag.Repository) CreateTagHandler {
	return CreateTagHandler{TagRepository: tagRepository}
}

func (h CreateTagHandler) Handle(tag tag.Tag) (*tag.Tag, error) {
	return h.TagRepository.CreateTag(tag)
}
