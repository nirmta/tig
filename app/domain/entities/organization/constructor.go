package organization

import (
	"tig/app/domain/entities/user"
)

func New() *Organization {
	return &Organization{
		users: make([]*user.User, 0),
	}
}

func Reconstitute(users []*user.User) *Organization {
	if users == nil {
		users = make([]*user.User, 0)
	}
	return &Organization{
		users: users,
	}
}
