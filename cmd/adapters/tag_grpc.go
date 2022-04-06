package adapters

import (
	"context"
	"encoding/json"
	"log"

	tagpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tag"
	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
	"google.golang.org/protobuf/encoding/protojson"
)

type TagGrpc struct {
	client tagpb.TagServiceClient
}

func NewTagGrpc(client tagpb.TagServiceClient) TagGrpc {
	return TagGrpc{client: client}
}

func (s TagGrpc) FindOrCreate(ctx context.Context, data tag.Tag) (*tag.Tag, error) {

	ee, err := json.Marshal(data)
	var tagPB tagpb.Tag
	um := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	err = um.Unmarshal(ee, &tagPB)
	log.Println(err)
	response, err := s.client.FindOrCreate(ctx, &tagpb.FindOrCreateRequest{
		Tag: &tagPB,
	})

	return &tag.Tag{Name: response.Data.Name}, err
}
