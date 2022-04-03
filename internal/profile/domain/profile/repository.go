package profile

type Repository interface {
	Find() ([]Profile, error)
	Create(profile Profile) error
	Update(profile Profile) error
}
