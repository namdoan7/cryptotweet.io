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

func (s TagGrpc) GetTag(ctx context.Context, name string) (*tag.Tag, error) {
	result, err := s.client.GetTag(ctx, &tagpb.GetTagRequest{})
	var tag *tag.Tag
	return tag.FromProtoMessage(result.Tag), err
}

func (s TagGrpc) CreateTag(ctx context.Context, tag tag.Tag) error {
	_, err := s.client.CreateTag(ctx, &tagpb.CreateTagRequest{
		Tag: tag.ToProtoMessage(),
	})
	return err
}

func (s TagGrpc) FindOrCreate(ctx context.Context, tag tag.Tag) error {
	_, err := s.client.FindOrCreate(ctx, &tagpb.FindOrCreateRequest{
		Tag: tag.ToProtoMessage(),
	})
	return err
}
