package usecases

import (
	"tig/application/repositories"
	"tig/application/services"
	UserEntity "tig/domain/entities/user"
)

type AddUserInput struct {
	Alias string
	Name  string
	Email string
	Host  string
}

type AddNewUserUseCaseInput struct {
	User AddUserInput
}

type AddNewUserUseCase struct {
	organizationRepository repositories.OrganizationRepository
	gitService             services.GitService
}

func NewAddNewUserUseCase(organizationRepository repositories.OrganizationRepository, gitService services.GitService) *AddNewUserUseCase {
	return &AddNewUserUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (anu *AddNewUserUseCase) Execute(input AddNewUserUseCaseInput) {
	anu.gitService.CreateSSHKey(input.User.Alias, input.User.Email)
	identityFile := anu.gitService.GetPrivateKeyPath(input.User.Alias)

	userDraf := UserEntity.UserDraft{
		Alias:        input.User.Alias,
		Name:         input.User.Name,
		Email:        input.User.Email,
		Host:         input.User.Host,
		IdentityFile: identityFile,
	}

	newUser := UserEntity.Constructor(userDraf)
	org, _ := anu.organizationRepository.GetUnique()
	org.AddUser(newUser)
	anu.organizationRepository.Update(org)
	anu.gitService.GenerateSSHConfig(org.GetUsers())
}
