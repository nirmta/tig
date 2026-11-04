package usecase

import (
	"fmt"
	"tig/app/application/repository"
	"tig/app/application/port"
	domainRepo "tig/app/domain/entities/repository"
	"tig/app/domain/entities/user"
)

type DownloadRepositoriesUseCase struct {
	organizationRepository repository.OrganizationRepository
	gitService             port.GitService
	fsService              port.FsService
}

func NewDownloadRepositoriesUseCase(
	organizationRepository repository.OrganizationRepository,
	gitService port.GitService,
	fsService port.FsService,
) *DownloadRepositoriesUseCase {
	return &DownloadRepositoriesUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
		fsService:              fsService,
	}
}

func (dr *DownloadRepositoriesUseCase) Execute() error {
	currentUser, err := dr.organizationRepository.GetCurrentUser()
	if err != nil || currentUser == nil {
		return user.ErrUserNotFound
	}

	repositories := currentUser.Repositories()
	if len(repositories) == 0 {
		return domainRepo.ErrRepositoriesNotFound
	}

	userPath := currentUser.Alias()
	if err := dr.fsService.CreateDirectory(userPath); err != nil {
		return fmt.Errorf("failed to create user directory: %w", err)
	}

	for _, repo := range repositories {
		repositoryPath := "./" + userPath + "/" + repo.Alias()
		if err := dr.gitService.CloneRepository(repo.Origin(), repositoryPath); err != nil {
			return fmt.Errorf("failed cloning repository %s: %w", repo.Alias(), err)
		}
	}

	return nil
}
