package usecases

import (
	"errors"
	"tig/app/application/repositories"
	"tig/app/application/services"
)

type CheckoutUserUseCase struct {
	organizationRepository repositories.OrganizationRepository
	gitService             services.GitService
}

func NewCheckoutUserUseCase(organizationRepository repositories.OrganizationRepository, gitService services.GitService) *CheckoutUserUseCase {
	return &CheckoutUserUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (cu *CheckoutUserUseCase) Execute(alias string) (bool, error) {
	user, err := cu.organizationRepository.GetUserByAlias(alias)
	if err != nil || user == nil {
		return false, errors.New("user not found")
	}

	updatedGlobalConfig := cu.gitService.UpdateGlobalGitConfig(services.GlobalGitConfigInput{
		Email: user.Email(),
		Name:  user.Name(),
		Alias: user.Alias(),
	})

	if !updatedGlobalConfig {
		return false, errors.New("fail updating global config")
	}

	return true, nil
}
