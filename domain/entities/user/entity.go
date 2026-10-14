package UserEntity

import "tig/domain/constants"

type User struct {
	alias        string
	name         string
	email        string
	host         string
	identityFile string
	inGlobal     bool
}

func (u *User) UpdateName(name string) {
	u.name = name
}

func (u *User) UpdateHost(host string) {
	u.host = host
}

func (u *User) UpdateIdentityFile(identityFile string) {
	u.identityFile = identityFile
}

func (u *User) UpdateInGlobal(inGlobal bool) {
	u.inGlobal = inGlobal
}

func (u *User) Alias() string {
	return u.alias
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Host() string {
	return u.host
}

func (u *User) RepositoryHost() string {
	return constants.RepositoryHosts[u.host]
}

func (u *User) IdentityFile() string {
	return u.identityFile
}

func (u *User) InGlobal() bool {
	return u.inGlobal
}
