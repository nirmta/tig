package usecases

import (
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

func (rr *RemoveRepositoryUseCase) Execute(alias string) {
	user, error := rr.organizationRepository.GetCurrentUser()

	if error != nil {
		return
	}

	user.RemoveRepositoryByAlias(alias)
	org, _ := rr.organizationRepository.GetUnique()
	org.UpdateUser(user)
	rr.organizationRepository.Update(org)
}
