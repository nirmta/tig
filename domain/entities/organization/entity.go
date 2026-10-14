package OrganizationEntity

import (
	UserEntity "tig/domain/entities/user"
)

type Organization struct {
	users []*UserEntity.User
}

func (o *Organization) SetUsers(users []*UserEntity.User) {
	o.users = users
}

func (o *Organization) AddUser(user *UserEntity.User) {
	foundUser := o.FindUserByAlias(user.Alias())

	if foundUser != nil {
		return
	}

	o.users = append(o.users, user)
}

func (o *Organization) RemoveUserByAlias(alias string) {
	for i, user := range o.users {
		if user.Alias() == alias {
			o.users = append(o.users[:i], o.users[i+1:]...)
			break
		}
	}
}

func (o *Organization) FindUserByAlias(alias string) *UserEntity.User {
	for _, user := range o.users {
		if user.Alias() == alias {
			return user
		}
	}

	return nil
}

func (o *Organization) UpdateUser(updatedUser *UserEntity.User) {
	for i, user := range o.users {
		if user.Alias() == updatedUser.Alias() {
			o.users[i] = updatedUser
			break
		}
	}
}

func (o *Organization) GetUsers() []*UserEntity.User {
	return o.users
}
