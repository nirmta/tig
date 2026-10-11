package usecases

import (
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

func (anr *AddNewRepositoryUseCase) Execute(input AddNewRepositoryUseCaseInput) {
	user, err := anr.organizationRepository.GetCurrentUser()

	if foundRepository := user.FindRepositoryByAlias(input.Alias); foundRepository != nil {
		return
	}

	if err != nil {
		return
	}

	if user == nil {
		return
	}

	newRepository := RepositoryEntity.Constructor(RepositoryEntity.RepositoryDraft{
		Alias:  input.Alias,
		Origin: input.Origin,
	})

	user.AddRepository(newRepository)

	org, err := anr.organizationRepository.GetUnique()

	if err != nil {
		return
	}

	org.UpdateUser(user)
	anr.organizationRepository.Update(org)
}
