package usecases

import (
	"tig/application/repositories"
	"tig/application/services"
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

func (cu *CheckoutUserUseCase) Execute(alias string) {
	user, err := cu.organizationRepository.GetUserByAlias(alias)

	if err != nil {
		return
	}

	cu.gitService.UpdateGlobalGitConfig(services.GlobalGitConfigInput{
		Email: user.Email(),
		Name:  user.Name(),
		Alias: user.Alias(),
	})
}
