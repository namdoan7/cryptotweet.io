package tag

type Repository interface {
	GetTag(name string) (*Tag, error)
	CreateTag(tag Tag) (*Tag, error)
	FindOrCreate(tag Tag) (*Tag, error)
}
