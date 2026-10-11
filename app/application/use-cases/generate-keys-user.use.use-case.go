package usecases

import (
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

func (gku *GenerateKeysUserUseCase) Execute() {
	user, _ := gku.organizationRepository.GetCurrentUser()

	if user == nil {
		return
	}

	gku.gitService.RemoveSSHKey(user.Alias())
	gku.gitService.CreateSSHKey(user.Alias(), user.Email())
}
