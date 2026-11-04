package repository

import "errors"

var (
	ErrEmptyAlias              = errors.New("repository alias cannot be empty")
	ErrEmptyOrigin             = errors.New("repository origin cannot be empty")
	ErrRepositoryAlreadyExists = errors.New("repository already exists")
	ErrRepositoriesNotFound    = errors.New("repositories not found")
)
