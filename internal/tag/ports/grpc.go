package ports

import (
	"context"
	"encoding/json"

	tagpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tag"
	"github.com/levinhne/cryptotweet.io/internal/tag/app"
	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func trans[R any](s protoreflect.ProtoMessage) R {
	m := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	var r R
	ee, _ := m.Marshal(s)
	json.Unmarshal(ee, &r)
	return r
}

type GrpcServer struct {
	app app.Application
	tagpb.UnimplementedTagServiceServer
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) GetTag(ctx context.Context, in *tagpb.GetTagRequest) (*tagpb.GetTagResponse, error) {
	return &tagpb.GetTagResponse{}, nil
}

func (g GrpcServer) FindOrCreate(ctx context.Context, in *tagpb.FindOrCreateRequest) (*tagpb.FindOrCreateResponse, error) {
	tag := trans[tag.Tag](in.Tag)
	data, err := g.app.Commands.FindOrCreate.Handle(tag)
	return &tagpb.FindOrCreateResponse{Data: &tagpb.Tag{Name: data.Name}}, err
}
