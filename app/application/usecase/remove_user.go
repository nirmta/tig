package usecase

import (
	"fmt"
	"tig/app/application/repository"
	"tig/app/application/port"
	"tig/app/domain/entities/organization"
	"tig/app/domain/entities/user"
)

type RemoveUserUseCase struct {
	organizationRepository repository.OrganizationRepository
	gitService             port.GitService
}

func NewRemoveUserUseCase(organizationRepository repository.OrganizationRepository, gitService port.GitService) *RemoveUserUseCase {
	return &RemoveUserUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (ru *RemoveUserUseCase) Execute(alias string) error {
	u, err := ru.organizationRepository.GetUserByAlias(alias)
	if err != nil || u == nil {
		return user.ErrUserNotFound
	}

	org, err := ru.organizationRepository.GetUnique()
	if err != nil || org == nil {
		return organization.ErrOrganizationNotFound
	}

	if err := ru.gitService.RemoveSSHKey(u.Alias()); err != nil {
		return fmt.Errorf("failed removing user ssh key: %w", err)
	}

	org.RemoveUserByAlias(alias)

	if err := ru.organizationRepository.Update(org); err != nil {
		return fmt.Errorf("failed removing user from organization: %w", err)
	}

	return nil
}
