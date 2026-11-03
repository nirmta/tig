package usecase

import (
	"fmt"
	"tig/app/application/repository"
	"tig/app/application/port"
	"tig/app/domain/entities/organization"
	"tig/app/domain/entities/user"
)

type RemoveRepositoryUseCase struct {
	organizationRepository repository.OrganizationRepository
	gitService             port.GitService
}

func NewRemoveRepositoryUseCase(organizationRepository repository.OrganizationRepository, gitService port.GitService) *RemoveRepositoryUseCase {
	return &RemoveRepositoryUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (rr *RemoveRepositoryUseCase) Execute(alias string) error {
	currentUser, err := rr.organizationRepository.GetCurrentUser()
	if err != nil || currentUser == nil {
		return user.ErrUserNotFound
	}

	currentUser.RemoveRepositoryByAlias(alias)

	org, err := rr.organizationRepository.GetUnique()
	if err != nil || org == nil {
		return organization.ErrOrganizationNotFound
	}

	org.UpdateUser(currentUser)

	if err := rr.organizationRepository.Update(org); err != nil {
		return fmt.Errorf("failed updating repositories: %w", err)
	}

	return nil
}
