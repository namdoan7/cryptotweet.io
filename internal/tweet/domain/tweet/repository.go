package tweet

type Repository interface {
	Create(tweet Tweet) error
	Update(tweet Tweet) error
}
