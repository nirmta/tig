package repository

import (
	"tig/app/domain/entities/organization"
	"tig/app/domain/entities/user"
)

type OrganizationRepository interface {
	GetUnique() (*organization.Organization, error)
	Update(org *organization.Organization) error
	GetUserByAlias(alias string) (*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	GetCurrentUser() (*user.User, error)
	GetAllUsers() ([]*user.User, error)
	CreateUser(user *user.User) error
	UpdateUser(user *user.User) error
}
