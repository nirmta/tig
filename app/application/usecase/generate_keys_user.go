package usecase

import (
	"fmt"
	"tig/app/application/repository"
	"tig/app/application/port"
	"tig/app/domain/entities/user"
)

type GenerateKeysUserUseCase struct {
	organizationRepository repository.OrganizationRepository
	gitService             port.GitService
}

func NewGenerateKeysUserUseCase(organizationRepository repository.OrganizationRepository, gitService port.GitService) *GenerateKeysUserUseCase {
	return &GenerateKeysUserUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (gku *GenerateKeysUserUseCase) Execute() error {
	currentUser, err := gku.organizationRepository.GetCurrentUser()
	if err != nil || currentUser == nil {
		return user.ErrUserNotFound
	}

	if err := gku.gitService.RemoveSSHKey(currentUser.Alias()); err != nil {
		return fmt.Errorf("failed removing ssh key: %w", err)
	}

	if err := gku.gitService.CreateSSHKey(currentUser.Alias(), currentUser.Email()); err != nil {
		return fmt.Errorf("failed creating ssh key: %w", err)
	}

	return nil
}
