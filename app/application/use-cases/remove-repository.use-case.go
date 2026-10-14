package usecases

import (
	"errors"
	"tig/app/application/repositories"
	"tig/app/application/services"
)

type RemoveRepositoryUseCase struct {
	organizationRepository repositories.OrganizationRepository
	gitService             services.GitService
}

func NewRemoveRepositoryUseCase(organizationRepository repositories.OrganizationRepository, gitService services.GitService) *RemoveRepositoryUseCase {
	return &RemoveRepositoryUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (rr *RemoveRepositoryUseCase) Execute(alias string) (bool, error) {
	user, err := rr.organizationRepository.GetCurrentUser()
	if user == nil || err != nil {
		return false, errors.New("current user not found")
	}

	user.RemoveRepositoryByAlias(alias)
	org, err := rr.organizationRepository.GetUnique()
	if err != nil || org == nil {
		return false, errors.New("organization not found")
	}

	org.UpdateUser(user)
	err = rr.organizationRepository.Update(org)
	if err != nil {
		return false, errors.New("fail updating repositories")
	}

	return false, nil
}
