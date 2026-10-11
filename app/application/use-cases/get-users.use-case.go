package usecases

import (
	"tig/app/application/repositories"
	"tig/app/application/services"
	UserEntity "tig/app/domain/entities/user"
)

type GetUsersUseCase struct {
	organizationRepository repositories.OrganizationRepository
	gitService             services.GitService
}

func NewGetUsersUseCase(organizationRepository repositories.OrganizationRepository, gitService services.GitService) *GetUsersUseCase {
	return &GetUsersUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (gu *GetUsersUseCase) Execute() []*UserEntity.User {
	org, _ := gu.organizationRepository.GetUnique()
	gitGlobalConfig, _ := gu.gitService.GetGlobalSettings()
	user, _ := gu.organizationRepository.GetUserByEmail(gitGlobalConfig.Email)

	if user != nil {
		user.UpdateInGlobal(true)
	}

	org.UpdateUser(user)
	return org.GetUsers()
}
