package OrganizationEntity

import (
	UserEntity "tig/domain/entities/user"
)

func Constructor() *Organization {
	return &Organization{
		users: []*UserEntity.User{},
	}
}
