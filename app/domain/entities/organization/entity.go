package organization

import (
	"tig/app/domain/entities/user"
)

type Organization struct {
	users []*user.User
}

func (o *Organization) SetUsers(users []*user.User) {
	o.users = users
}

func (o *Organization) AddUser(u *user.User) {
	foundUser := o.FindUserByAlias(u.Alias())

	if foundUser != nil {
		return
	}

	o.users = append(o.users, u)
}

func (o *Organization) RemoveUserByAlias(alias string) {
	for i, u := range o.users {
		if u.Alias() == alias {
			o.users = append(o.users[:i], o.users[i+1:]...)
			break
		}
	}
}

func (o *Organization) FindUserByAlias(alias string) *user.User {
	for _, u := range o.users {
		if u.Alias() == alias {
			return u
		}
	}

	return nil
}

func (o *Organization) UpdateUser(updatedUser *user.User) {
	for i, u := range o.users {
		if u.Alias() == updatedUser.Alias() {
			o.users[i] = updatedUser
			break
		}
	}
}

func (o *Organization) GetUsers() []*user.User {
	return o.users
}
