package OrganizationEntity

import (
	UserEntity "tig/app/domain/entities/user"
)

func Constructor() *Organization {
	return &Organization{
		users: []*UserEntity.User{},
	}
}
