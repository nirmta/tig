package usecase

import (
	appRepo "tig/app/application/repository"
	domainRepo "tig/app/domain/entities/repository"
	"tig/app/domain/entities/user"
)

type GetRepositoriesUseCase struct {
	organizationRepository appRepo.OrganizationRepository
}

func NewGetRepositoriesUseCase(organizationRepository appRepo.OrganizationRepository) *GetRepositoriesUseCase {
	return &GetRepositoriesUseCase{
		organizationRepository: organizationRepository,
	}
}

func (gr *GetRepositoriesUseCase) Execute() ([]*domainRepo.Repository, error) {
	currentUser, err := gr.organizationRepository.GetCurrentUser()
	if currentUser == nil || err != nil {
		return []*domainRepo.Repository{}, user.ErrUserNotFound
	}

	return currentUser.Repositories(), nil
}
