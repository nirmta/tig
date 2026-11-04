package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"tig/app/application/port"
	"tig/app/infrastructure/consts"
)

type BackupUseCaseInput struct {
	DestinationPath string
}

type BackupUseCase struct {
	tomlService port.TomlService
}

func NewBackupUseCase(tomlService port.TomlService) *BackupUseCase {
	return &BackupUseCase{
		tomlService: tomlService,
	}
}

func (bu *BackupUseCase) Execute(input BackupUseCaseInput) error {
	destPath := input.DestinationPath
	if destPath == "" {
		destPath = "organization-backup.toml"
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	systemFolder := filepath.Join(homeDir, consts.FOLDER)
	data, err := bu.tomlService.ReadFile(systemFolder, consts.FILE_NAME)
	if err != nil {
		return fmt.Errorf("failed reading organization configuration: %w", err)
	}

	if err := os.WriteFile(destPath, data, 0644); err != nil {
		return fmt.Errorf("failed writing backup file: %w", err)
	}

	return nil
}
