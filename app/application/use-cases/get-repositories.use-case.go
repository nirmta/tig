package usecases

import (
	"errors"
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

func (gr *GetRepositoriesUseCase) Execute() ([]*RepositoryEntity.Repository, error) {
	user, err := gr.organizationRepository.GetCurrentUser()

	if user == nil || err != nil {
		return []*RepositoryEntity.Repository{}, errors.New("user not found")
	}

	repositories := user.Repositories()
	return repositories, nil
}
