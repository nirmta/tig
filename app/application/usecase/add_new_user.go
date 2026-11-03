package usecase

import (
	"fmt"
	"tig/app/application/repository"
	"tig/app/application/port"
	"tig/app/domain/entities/organization"
	"tig/app/domain/entities/user"
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
	organizationRepository repository.OrganizationRepository
	gitService             port.GitService
}

func NewAddNewUserUseCase(organizationRepository repository.OrganizationRepository, gitService port.GitService) *AddNewUserUseCase {
	return &AddNewUserUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (anu *AddNewUserUseCase) Execute(input AddNewUserUseCaseInput) error {
	if err := anu.gitService.CreateSSHKey(input.User.Alias, input.User.Email); err != nil {
		return fmt.Errorf("failed creating ssh key: %w", err)
	}

	identityFile := anu.gitService.GetPrivateKeyPath(input.User.Alias)

	userDraft := user.Draft{
		Alias:        input.User.Alias,
		Name:         input.User.Name,
		Email:        input.User.Email,
		Host:         input.User.Host,
		IdentityFile: identityFile,
	}

	newUser, err := user.New(userDraft)
	if err != nil {
		return err
	}

	org, err := anu.organizationRepository.GetUnique()
	if err != nil || org == nil {
		return organization.ErrOrganizationNotFound
	}

	org.AddUser(newUser)

	if err := anu.organizationRepository.Update(org); err != nil {
		return fmt.Errorf("failed updating organization: %w", err)
	}

	if err := anu.gitService.GenerateSSHConfig(org.GetUsers()); err != nil {
		return fmt.Errorf("failed setting up ssh config: %w", err)
	}

	return nil
}
