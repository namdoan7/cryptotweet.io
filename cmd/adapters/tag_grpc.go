package adapters

import (
	"context"

	tagpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tag"
	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
)

type TagGrpc struct {
	client tagpb.TagServiceClient
}

func NewTagGrpc(client tagpb.TagServiceClient) TagGrpc {
	return TagGrpc{client: client}
}

func (s TagGrpc) FindOrCreate(ctx context.Context, name string) (tag.Tag, error) {
	response, err := s.client.FindOrCreate(ctx, &tagpb.FindOrCreateRequest{
		Name: name,
	})
	return tag.Tag{Name: response.Data.Name}, err
}
