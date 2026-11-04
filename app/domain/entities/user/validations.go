package user

import (
	"net/mail"
	"strings"
	"tig/app/domain/constants"
)

func ValidateDraft(d Draft) error {
	if strings.TrimSpace(d.Alias) == "" {
		return ErrEmptyAlias
	}
	if strings.TrimSpace(d.Name) == "" {
		return ErrEmptyName
	}
	if _, err := mail.ParseAddress(d.Email); err != nil {
		return ErrInvalidEmail
	}
	if _, ok := constants.RepositoryHosts[d.Host]; !ok && d.Host != "" {
		return ErrInvalidHost
	}
	return nil
}
