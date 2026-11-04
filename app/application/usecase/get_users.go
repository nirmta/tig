package usecase

import (
	"tig/app/application/repository"
	"tig/app/application/port"
	"tig/app/domain/entities/organization"
	"tig/app/domain/entities/user"
)

type GetUsersUseCase struct {
	organizationRepository repository.OrganizationRepository
	gitService             port.GitService
}

func NewGetUsersUseCase(organizationRepository repository.OrganizationRepository, gitService port.GitService) *GetUsersUseCase {
	return &GetUsersUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (gu *GetUsersUseCase) Execute() ([]*user.User, error) {
	org, err := gu.organizationRepository.GetUnique()
	if org == nil || err != nil {
		return []*user.User{}, organization.ErrOrganizationNotFound
	}

	gitGlobalConfig, _ := gu.gitService.GetGlobalSettings()
	u, _ := gu.organizationRepository.GetUserByEmail(gitGlobalConfig.Email)

	if u != nil {
		u.UpdateInGlobal(true)
		org.UpdateUser(u)
	}

	return org.GetUsers(), nil
}
