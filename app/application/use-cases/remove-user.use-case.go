package usecases

import (
	"errors"
	"tig/app/application/repositories"
	"tig/app/application/services"
)

type RemoveUserUseCase struct {
	organizationRepository repositories.OrganizationRepository
	gitService             services.GitService
}

func NewRemoveUserUseCase(organizationRepository repositories.OrganizationRepository, gitService services.GitService) *RemoveUserUseCase {
	return &RemoveUserUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (ru *RemoveUserUseCase) Execute(alias string) (bool, error) {
	user, err := ru.organizationRepository.GetUserByAlias(alias)
	if err != nil || user == nil {
		return false, errors.New("user not found")
	}

	org, err := ru.organizationRepository.GetUnique()
	if err != nil || org == nil {
		return false, errors.New("organization not found")
	}

	removedSSHKey := ru.gitService.RemoveSSHKey(user.Alias())
	if !removedSSHKey {
		return false, errors.New("fail removing user ssh key")
	}

	org.RemoveUserByAlias(user.Alias())
	err = ru.organizationRepository.Update(org)
	if err != nil {
		return false, errors.New("fail removing user")
	}

	return true, nil
}
