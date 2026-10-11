package UserEntity

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrEmptyName    = errors.New("name cannot be empty")
)
