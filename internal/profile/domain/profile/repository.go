package profile

type Repository interface {
	Create(profile Profile) error
	Update(profile Profile) error
}
