package usecases

import (
	"errors"
	"tig/app/application/repositories"
	"tig/app/application/services"
)

type GenerateKeysUserUseCase struct {
	organizationRepository repositories.OrganizationRepository
	gitService             services.GitService
}

func NewGenerateKeysUserUseCase(organizationRepository repositories.OrganizationRepository, gitService services.GitService) *GenerateKeysUserUseCase {
	return &GenerateKeysUserUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (gku *GenerateKeysUserUseCase) Execute() (bool, error) {
	user, err := gku.organizationRepository.GetCurrentUser()
	if err != nil || user == nil {
		return false, errors.New("user not found")
	}

	removedSSHKey := gku.gitService.RemoveSSHKey(user.Alias())
	if !removedSSHKey {
		return false, errors.New("fail removing ssh key")
	}

	createdSSHKey := gku.gitService.CreateSSHKey(user.Alias(), user.Email())
	if !createdSSHKey {
		return false, errors.New("fail created ssh key")
	}

	return true, nil
}
