package RepositoryEntity

type RepositoryDraft struct {
	Alias  string
	Origin string
}

func Constructor(userDraft RepositoryDraft) *Repository {
	return &Repository{
		alias:  userDraft.Alias,
		origin: userDraft.Origin,
	}
}
