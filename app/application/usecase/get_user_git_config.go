package usecase

import (
	"tig/app/application/repository"
	"tig/app/application/port"
	"tig/app/domain/entities/user"
)

type GetUserGitConfigUseCaseOutput struct {
	GlobalSettings port.GlobalGitConfigOutput
	SSHConfig      port.UserSSHConfigOutput
}

type GetUserGitConfigUseCase struct {
	organizationRepository repository.OrganizationRepository
	gitService             port.GitService
}

func NewGetUserGitConfigUseCase(organizationRepository repository.OrganizationRepository, gitService port.GitService) *GetUserGitConfigUseCase {
	return &GetUserGitConfigUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (gupk *GetUserGitConfigUseCase) Execute() (GetUserGitConfigUseCaseOutput, error) {
	output := GetUserGitConfigUseCaseOutput{}

	currentUser, err := gupk.organizationRepository.GetCurrentUser()
	if err != nil || currentUser == nil {
		return output, user.ErrUserNotFound
	}

	userSSHConfig, err := gupk.gitService.GetUserSSHConfig(currentUser.Alias())
	if err != nil {
		return output, user.ErrUserSSHConfigNotFound
	}

	globalSettings, err := gupk.gitService.GetGlobalSettings()
	if err != nil {
		return output, err
	}

	output.SSHConfig = userSSHConfig
	output.GlobalSettings = globalSettings

	return output, nil
}
