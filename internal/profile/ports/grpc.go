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

func (g GrpcServer) Find(ctx context.Context, in *profilepb.FindProfileRequest) (*profilepb.FindProfileResponse, error) {
	proflies, _ := g.app.Queries.FindProfile.Handle()
	data := make([]*profilepb.FindProfileResponse_Profile, 0)
	for _, profile := range proflies {
		data = append(data, &profilepb.FindProfileResponse_Profile{
			ScreenName: profile.ScreenName,
		})
	}

	return &profilepb.FindProfileResponse{Data: data}, nil
}

func (g GrpcServer) Create(ctx context.Context, in *profilepb.CreateProfileRequest) (*profilepb.CreateProfileResponse, error) {
	profile := trans[profile.Profile](in)
	g.app.Commands.CreateProfile.Handle(profile)
	return &profilepb.CreateProfileResponse{}, nil
}
