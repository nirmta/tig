package usecases

import (
	"errors"
	"tig/app/application/repositories"
	"tig/app/application/services"
)

type DownloadRepositoriesUseCase struct {
	organizationRepository repositories.OrganizationRepository
	gitService             services.GitService
	fsService              services.FsService
}

func NewDownloadRepositoriesUseCase(
	organizationRepository repositories.OrganizationRepository,
	gitService services.GitService,
	fsService services.FsService,
) *DownloadRepositoriesUseCase {
	return &DownloadRepositoriesUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
		fsService:              fsService,
	}
}

func (dr *DownloadRepositoriesUseCase) Execute() (bool, error) {
	user, err := dr.organizationRepository.GetCurrentUser()
	if err != nil || user == nil {
		return false, errors.New("user not found")
	}

	repositories := user.Repositories()
	if len(repositories) == 0 {
		return false, errors.New("repositories not found")
	}

	userPath := user.Alias() + "-repositories"
	dr.fsService.CreateDirectory(userPath)

	for _, repository := range repositories {
		repositoryPath := "./" + userPath + "/" + repository.Alias()
		dr.gitService.CloneRepository(repository.Origin(), repositoryPath)
	}

	return true, nil
}
