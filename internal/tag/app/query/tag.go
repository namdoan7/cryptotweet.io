package query

import (
	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
)

type GetTagHandler struct {
	TagRepository tag.Repository
}

func NewGetTagHandler(tagRepository tag.Repository) *GetTagHandler {
	return &GetTagHandler{TagRepository: tagRepository}
}

func (h GetTagHandler) Handle(name string) (*tag.Tag, error) {
	return h.TagRepository.GetTag(name)
}
