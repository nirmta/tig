package repositories

import (
	OrganizationEntity "tig/domain/entities/organization"
	UserEntity "tig/domain/entities/user"
)

type OrganizationRepository interface {
	GetUnique() (*OrganizationEntity.Organization, error)
	Update(org *OrganizationEntity.Organization) error
	GetUserByAlias(alias string) (*UserEntity.User, error)
	GetUserByEmail(email string) (*UserEntity.User, error)
	GetAllUsers() ([]*UserEntity.User, error)
	CreateUser(user *UserEntity.User) error
	UpdateUser(user *UserEntity.User) error
}
