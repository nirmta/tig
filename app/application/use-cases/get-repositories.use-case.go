package usecases

import (
	"tig/app/application/repositories"
	RepositoryEntity "tig/app/domain/entities/repository"
)

type GetRepositoriesUseCase struct {
	organizationRepository repositories.OrganizationRepository
}

func NewGetRepositoriesUseCase(organizationRepository repositories.OrganizationRepository) *GetRepositoriesUseCase {
	return &GetRepositoriesUseCase{
		organizationRepository: organizationRepository,
	}
}

func (gr *GetRepositoriesUseCase) Execute() []*RepositoryEntity.Repository {
	user, _ := gr.organizationRepository.GetCurrentUser()

	if user == nil {
		return []*RepositoryEntity.Repository{}
	}

	repositories := user.Repositories()

	return repositories
}
