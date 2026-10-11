package usecases

import (
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

func (dr *DownloadRepositoriesUseCase) Execute() {
	user, _ := dr.organizationRepository.GetCurrentUser()

	if user == nil {
		return
	}

	repositories := user.Repositories()

	userPath := user.Alias() + "-repositories"
	if len(repositories) > 0 {
		dr.fsService.CreateDirectory(userPath)
	}

	for _, repository := range repositories {
		repositoryPath := "./" + userPath + "/" + repository.Alias()
		dr.gitService.CloneRepository(repository.Origin(), repositoryPath)
	}
}
