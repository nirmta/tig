package user

import "errors"

var (
	ErrInvalidEmail       = errors.New("invalid email")
	ErrEmptyName          = errors.New("name cannot be empty")
	ErrEmptyAlias         = errors.New("alias cannot be empty")
	ErrInvalidHost        = errors.New("host is not supported")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserSSHConfigNotFound = errors.New("user ssh config not found")
)
