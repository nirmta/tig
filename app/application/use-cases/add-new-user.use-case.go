package usecases

import (
	"errors"
	"tig/app/application/repositories"
	"tig/app/application/services"
	UserEntity "tig/app/domain/entities/user"
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

func (anu *AddNewUserUseCase) Execute(input AddNewUserUseCaseInput) (bool, error) {
	createdShhKey := anu.gitService.CreateSSHKey(input.User.Alias, input.User.Email)
	if !createdShhKey {
		return false, errors.New("fail creating ssh key")
	}

	identityFile := anu.gitService.GetPrivateKeyPath(input.User.Alias)

	userDraf := UserEntity.UserDraft{
		Alias:        input.User.Alias,
		Name:         input.User.Name,
		Email:        input.User.Email,
		Host:         input.User.Host,
		IdentityFile: identityFile,
	}

	newUser := UserEntity.Constructor(userDraf)
	org, err := anu.organizationRepository.GetUnique()
	if err != nil || org == nil {
		return false, errors.New("organization not found")
	}

	org.AddUser(newUser)

	err = anu.organizationRepository.Update(org)
	if err != nil {
		return false, errors.New("fail updating organization")
	}

	generatedSSHConfig := anu.gitService.GenerateSSHConfig(org.GetUsers())
	if !generatedSSHConfig {
		return false, errors.New("fail creating setting up ssh config")
	}

	return true, nil
}
