package adapters

import (
	"context"

	tagpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tag"
)

type TagGrpc struct {
	client tagpb.TagServiceClient
}

func NewTagGrpc(client tagpb.TagServiceClient) TagGrpc {
	return TagGrpc{client: client}
}

func (s TagGrpc) FindOrCreate(ctx context.Context, name string) error {
	return nil
}
