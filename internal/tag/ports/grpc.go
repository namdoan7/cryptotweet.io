package ports

import (
	"context"
	"encoding/json"

	tagpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tag"
	"github.com/levinhne/cryptotweet.io/internal/tag/app"
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

func (g GrpcServer) FindOrCreate(ctx context.Context, in *tagpb.FindOrCreateRequest) (*tagpb.FindOrCreateResponse, error) {
	tag, err := g.app.Commands.FindOrCreate.Handle(in.Name)
	return &tagpb.FindOrCreateResponse{Data: &tagpb.Tag{Name: tag.Name}}, err
}
