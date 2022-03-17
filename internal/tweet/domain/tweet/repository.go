package tweet

type Repository interface {
	Create(tweet Tweet) error
}
