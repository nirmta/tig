package usecases

import (
	"errors"
	"tig/app/application/repositories"
	"tig/app/application/services"
	RepositoryEntity "tig/app/domain/entities/repository"
)

type AddNewRepositoryUseCaseInput struct {
	Alias  string
	Origin string
}

type AddNewRepositoryUseCase struct {
	organizationRepository repositories.OrganizationRepository
	gitService             services.GitService
}

func NewAddNewRepositoryUseCase(organizationRepository repositories.OrganizationRepository, gitService services.GitService) *AddNewRepositoryUseCase {
	return &AddNewRepositoryUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (anr *AddNewRepositoryUseCase) Execute(input AddNewRepositoryUseCaseInput) (bool, error) {
	user, err := anr.organizationRepository.GetCurrentUser()

	if foundRepository := user.FindRepositoryByAlias(input.Alias); foundRepository != nil {
		return false, errors.New("repository already exists")
	}

	if err != nil || user == nil {
		return false, errors.New("user not found")
	}

	newRepository := RepositoryEntity.Constructor(RepositoryEntity.RepositoryDraft{
		Alias:  input.Alias,
		Origin: input.Origin,
	})

	user.AddRepository(newRepository)

	org, err := anr.organizationRepository.GetUnique()
	if err != nil {
		return false, errors.New("organization not found")
	}

	org.UpdateUser(user)
	err = anr.organizationRepository.Update(org)

	if err != nil {
		return false, errors.New("error updating organization")
	}

	return true, nil
}
