package RepositoryEntity

type Repository struct {
	alias  string
	origin string
}

func (r *Repository) Alias() string {
	return r.alias
}

func (r *Repository) Origin() string {
	return r.origin
}
