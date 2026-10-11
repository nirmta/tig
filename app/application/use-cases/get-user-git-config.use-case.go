package usecases

import (
	"tig/app/application/repositories"
	"tig/app/application/services"
)

type GetUserGitConfigUseCase struct {
	organizationRepository repositories.OrganizationRepository
	gitService             services.GitService
}

func NewGetUserGitConfigUseCase(organizationRepository repositories.OrganizationRepository, gitService services.GitService) *GetUserGitConfigUseCase {
	return &GetUserGitConfigUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

type GetUserGitConfigUseCaseOutput struct {
	PublicKey string
}

func (gupk *GetUserGitConfigUseCase) Execute() (GetUserGitConfigUseCaseOutput, error) {
	currentUser, _ := gupk.organizationRepository.GetCurrentUser()
	var getUserGitConfigUseCaseOutput GetUserGitConfigUseCaseOutput
	user, err := gupk.organizationRepository.GetUserByAlias(currentUser.Alias())

	if user == nil {
		return getUserGitConfigUseCaseOutput, err
	}

	userSSHConfig, _ := gupk.gitService.GetUserSSHConfig(user.Alias())

	getUserGitConfigUseCaseOutput.PublicKey = userSSHConfig.PublicKey
	return getUserGitConfigUseCaseOutput, nil
}
