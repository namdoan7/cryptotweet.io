package command

import (
	"log"

	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
)

type FindOrCreateTagHandler struct {
	TagRepository tag.Repository
}

func NewFindOrCreateTagHandler(tagRepository tag.Repository) *FindOrCreateTagHandler {
	return &FindOrCreateTagHandler{TagRepository: tagRepository}
}

func (h FindOrCreateTagHandler) Handle(tag tag.Tag) (*tag.Tag, error) {
	log.Println(tag)
	return h.TagRepository.FindOrCreate(tag)
}
