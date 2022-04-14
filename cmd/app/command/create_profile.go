package command

import (
	"context"

	"github.com/levinhne/cryptotweet.io/cmd/adapters"
	"github.com/levinhne/cryptotweet.io/internal/profile/domain/profile"
)

type CreateProfile struct {
	Profile *profile.Profile
}

type CreateProfileHandler struct {
	ProfileService adapters.ProfileService
}

func NewCreateProfileHandler(profileService adapters.ProfileService) CreateProfileHandler {
	return CreateProfileHandler{ProfileService: profileService}
}

func (h CreateProfileHandler) Handle(ctx context.Context, cmd CreateProfile) error {
	h.ProfileService.CreateProfile(ctx, cmd.Profile)
	return nil
}
