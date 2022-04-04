package tag

type Repository interface {
	FindOrCreate(name string) (Tag, error)
}
