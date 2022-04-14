package query

import "github.com/levinhne/cryptotweet.io/internal/profile/domain/profile"

type FindProfileHandler struct {
	ProfileRepository profile.Repository
}

func NewFindProfileHandler(profileRepository profile.Repository) FindProfileHandler {
	return FindProfileHandler{ProfileRepository: profileRepository}
}

func (h FindProfileHandler) Handle() ([]profile.Profile, error) {
	return h.ProfileRepository.Find()
}
