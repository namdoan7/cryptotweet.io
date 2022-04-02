package ports

import (
	"context"
	"encoding/json"

	profilepb "github.com/levinhne/cryptotweet.io/internal/common/genproto/profile"
	"github.com/levinhne/cryptotweet.io/internal/profile/app"
	"github.com/levinhne/cryptotweet.io/internal/profile/domain/profile"
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
	profilepb.UnimplementedProfileServiceServer
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) Create(ctx context.Context, in *profilepb.CreateProfileRequest) (*profilepb.CreateProfileResponse, error) {
	profile := trans[profile.Profile](in)
	g.app.Commands.CreateProfile.Handle(profile)
	return &profilepb.CreateProfileResponse{}, nil
}