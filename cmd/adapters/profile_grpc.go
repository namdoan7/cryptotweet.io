package adapters

import (
	"context"
	"encoding/json"

	profilepb "github.com/levinhne/cryptotweet.io/internal/common/genproto/profile"
	profile "github.com/levinhne/cryptotweet.io/internal/profile/domain/profile"
	"google.golang.org/protobuf/encoding/protojson"
)

type ProfileGrpc struct {
	client profilepb.ProfileServiceClient
}

func NewProfileGrpc(client profilepb.ProfileServiceClient) ProfileGrpc {
	return ProfileGrpc{client: client}
}

func (s ProfileGrpc) CreateProfile(ctx context.Context, proflile *profile.Profile) error {
	ee, err := json.Marshal(proflile)
	var profileP profilepb.Profile
	um := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	err = um.Unmarshal(ee, &profileP)
	_, err = s.client.Create(ctx, &profilepb.CreateProfileRequest{Profile: &profileP})
	return err
}

// func (s ProfileGrpc) UpdateTweet(ctx context.Context, tweet *profile.Profile) error {
// 	ee, err := json.Marshal(proflile)

// 	var updateTweetRequest profilepb.UpdateTweetRequest
// 	um := protojson.UnmarshalOptions{
// 		DiscardUnknown: true,
// 	}
// 	err = um.Unmarshal(ee, &updateTweetRequest)
// 	log.Println(111111, updateTweetRequest.String())
// 	_, err = s.client.Update(ctx, &updateTweetRequest)
// 	return err
// }
