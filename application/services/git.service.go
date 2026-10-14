package services

import UserEntity "tig/domain/entities/user"

type GlobalGitConfigInput struct {
	Name  string
	Email string
	Alias string
}

type GlobalGitConfigOutput struct {
	Name  string
	Email string
	Alias string
}

type UserSSHConfigOutput struct {
	PublicKey  string
	PrivateKey string
}

type GitService interface {
	UpdateGlobalGitConfig(config GlobalGitConfigInput) bool
	CreateSSHKey(alias, email string) bool
	GetPrivateKeyPath(alias string) string
	RemoveSSHKey(alias string) bool
	GenerateSSHConfig(users []*UserEntity.User) bool
	GetGlobalSettings() (GlobalGitConfigOutput, error)
	GetUserSSHConfig(alias string) (UserSSHConfigOutput, error)
}
