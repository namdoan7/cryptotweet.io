package tag

type Repository interface {
	FindOrCreate(tag Tag) (*Tag, error)
}
