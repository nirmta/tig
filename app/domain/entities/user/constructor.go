package user

import (
	"tig/app/domain/entities/repository"
)

type Draft struct {
	Alias        string
	Name         string
	Email        string
	Host         string
	IdentityFile string
}

func New(draft Draft) (*User, error) {
	if err := ValidateDraft(draft); err != nil {
		return nil, err
	}

	return &User{
		alias:        draft.Alias,
		name:         draft.Name,
		email:        draft.Email,
		host:         draft.Host,
		identityFile: draft.IdentityFile,
		inGlobal:     false,
		repositories: make([]*repository.Repository, 0),
	}, nil
}

func Reconstitute(draft Draft, inGlobal bool, repos []*repository.Repository) *User {
	if repos == nil {
		repos = make([]*repository.Repository, 0)
	}

	return &User{
		alias:        draft.Alias,
		name:         draft.Name,
		email:        draft.Email,
		host:         draft.Host,
		identityFile: draft.IdentityFile,
		inGlobal:     inGlobal,
		repositories: repos,
	}
}
