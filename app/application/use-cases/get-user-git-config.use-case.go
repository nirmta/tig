package usecases

import (
	"errors"
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
	user, err := gupk.organizationRepository.GetCurrentUser()
	var getUserGitConfigUseCaseOutput GetUserGitConfigUseCaseOutput
	if err != nil || user == nil {
		return getUserGitConfigUseCaseOutput, errors.New("user not found")
	}

	userSSHConfig, err := gupk.gitService.GetUserSSHConfig(user.Alias())
	if err != nil {
		return getUserGitConfigUseCaseOutput, errors.New("user ssh config not found")
	}

	getUserGitConfigUseCaseOutput.PublicKey = userSSHConfig.PublicKey
	return getUserGitConfigUseCaseOutput, nil
}
