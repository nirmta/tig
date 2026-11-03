package usecase

import (
	"fmt"
	"tig/app/application/repository"
	"tig/app/application/port"
	"tig/app/domain/entities/user"
)

type CheckoutUserUseCase struct {
	organizationRepository repository.OrganizationRepository
	gitService             port.GitService
}

func NewCheckoutUserUseCase(organizationRepository repository.OrganizationRepository, gitService port.GitService) *CheckoutUserUseCase {
	return &CheckoutUserUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (cu *CheckoutUserUseCase) Execute(alias string) error {
	u, err := cu.organizationRepository.GetUserByAlias(alias)
	if err != nil || u == nil {
		return user.ErrUserNotFound
	}

	if err := cu.gitService.UpdateGlobalGitConfig(port.GlobalGitConfigInput{
		Email: u.Email(),
		Name:  u.Name(),
		Alias: u.Alias(),
	}); err != nil {
		return fmt.Errorf("failed updating global config: %w", err)
	}

	return nil
}
