package user

import (
	"tig/app/domain/constants"
	"tig/app/domain/entities/repository"
)

type User struct {
	alias        string
	name         string
	email        string
	host         string
	identityFile string
	inGlobal     bool
	repositories []*repository.Repository
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

func (u *User) Repositories() []*repository.Repository {
	return u.repositories
}

func (u *User) AddRepository(newRepository *repository.Repository) {
	u.repositories = append(u.repositories, newRepository)
}

func (u *User) FindRepositoryByAlias(alias string) *repository.Repository {
	for _, repo := range u.repositories {
		if repo.Alias() == alias {
			return repo
		}
	}

	return nil
}

func (u *User) RemoveRepositoryByAlias(repositoryAlias string) {
	var filteredRepositories []*repository.Repository

	for _, repo := range u.repositories {
		if repo.Alias() != repositoryAlias {
			filteredRepositories = append(filteredRepositories, repo)
		}
	}

	u.repositories = filteredRepositories
}
