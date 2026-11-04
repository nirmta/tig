package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"tig/app/application/port"
	"tig/app/application/repository"
	"tig/app/infrastructure/consts"
)

type RestoreUseCaseInput struct {
	SourcePath string
}

type RestoreUseCase struct {
	tomlService port.TomlService
	orgRepo     repository.OrganizationRepository
	gitService  port.GitService
}

func NewRestoreUseCase(
	tomlService port.TomlService,
	orgRepo repository.OrganizationRepository,
	gitService port.GitService,
) *RestoreUseCase {
	return &RestoreUseCase{
		tomlService: tomlService,
		orgRepo:     orgRepo,
		gitService:  gitService,
	}
}

func (ru *RestoreUseCase) Execute(input RestoreUseCaseInput) error {
	data, err := os.ReadFile(input.SourcePath)
	if err != nil {
		return fmt.Errorf("failed reading backup file: %w", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	systemFolder := filepath.Join(homeDir, consts.FOLDER)
	destPath := filepath.Join(systemFolder, ru.tomlService.AddExtension(consts.FILE_NAME))
	if err := os.WriteFile(destPath, data, 0600); err != nil {
		return fmt.Errorf("failed writing restored configuration: %w", err)
	}

	users, err := ru.orgRepo.GetAllUsers()
	if err != nil {
		return fmt.Errorf("failed reading restored users: %w", err)
	}

	if err := ru.gitService.GenerateSSHConfig(users); err != nil {
		return fmt.Errorf("failed regenerating ssh config: %w", err)
	}

	return nil
}
