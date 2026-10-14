package usecases

import (
	"tig/application/repositories"
	"tig/application/services"
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

func (ru *RemoveUserUseCase) Execute(alias string) error {
	user, err := ru.organizationRepository.GetUserByAlias(alias)

	if err != nil {
		return err
	}

	org, _ := ru.organizationRepository.GetUnique()

	ru.gitService.RemoveSSHKey(user.Alias())
	org.RemoveUserByAlias(user.Alias())
	ru.organizationRepository.Update(org)

	return nil
}
