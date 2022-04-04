package command

import "github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"

type FindOrCreateTagHandler struct {
	TagRepository tag.Repository
}

func NewFindOrCreateTagHandler(tagRepository tag.Repository) *FindOrCreateTagHandler {
	return &FindOrCreateTagHandler{TagRepository: tagRepository}
}

func (h FindOrCreateTagHandler) Handle(name string) (tag.Tag, error) {
	return h.TagRepository.FindOrCreate(name)
}
