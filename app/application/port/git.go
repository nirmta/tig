package port

import "tig/app/domain/entities/user"

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
	UpdateGlobalGitConfig(config GlobalGitConfigInput) error
	CreateSSHKey(alias, email string) error
	GetPrivateKeyPath(alias string) string
	RemoveSSHKey(alias string) error
	GenerateSSHConfig(users []*user.User) error
	GetGlobalSettings() (GlobalGitConfigOutput, error)
	GetUserSSHConfig(alias string) (UserSSHConfigOutput, error)
	CloneRepository(origin string, path string) error
}
