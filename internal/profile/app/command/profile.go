package command

import "github.com/levinhne/cryptotweet.io/internal/profile/domain/profile"

type CreateProfileHandler struct {
	ProfileRepository profile.Repository
}

func NewCreateProfileHandler(profileRepository profile.Repository) CreateProfileHandler {
	return CreateProfileHandler{ProfileRepository: profileRepository}
}

func (h CreateProfileHandler) Handle(profile profile.Profile) error {
	return h.ProfileRepository.Create(profile)
}
