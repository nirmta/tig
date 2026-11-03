package repository

import (
	"strings"
)

type Draft struct {
	Alias  string
	Origin string
}

func New(draft Draft) (*Repository, error) {
	if strings.TrimSpace(draft.Alias) == "" {
		return nil, ErrEmptyAlias
	}
	if strings.TrimSpace(draft.Origin) == "" {
		return nil, ErrEmptyOrigin
	}

	return &Repository{
		alias:  draft.Alias,
		origin: draft.Origin,
	}, nil
}

func Reconstitute(draft Draft) *Repository {
	return &Repository{
		alias:  draft.Alias,
		origin: draft.Origin,
	}
}
