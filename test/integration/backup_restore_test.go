package integration_test

import (
	"os"
	"path/filepath"
	"testing"
	"tig/app/application/usecase"
	"tig/app/infrastructure/consts"
	"tig/app/infrastructure/persistence/repositories"
	"tig/app/infrastructure/services"
)

func TestIntegration_Backup_And_Restore_UseCases(t *testing.T) {
	// Arrange
	tempDir := t.TempDir()
	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	t.Cleanup(func() {
		os.Setenv("HOME", origHome)
	})

	tomlService := services.NewTomlService()
	fsService := services.NewFsService()
	gitService := services.NewGitService(fsService)
	orgRepo := repositories.NewOrganizationRepository(tomlService, fsService, gitService)

	systemFolder := fsService.GetFolderPath([]string{consts.FOLDER})
	if err := fsService.CreateDirectory(systemFolder); err != nil {
		t.Fatalf("failed creating system folder: %v", err)
	}

	addNewUserUseCase := usecase.NewAddNewUserUseCase(orgRepo, gitService)
	addNewRepoUseCase := usecase.NewAddNewRepositoryUseCase(orgRepo, gitService)

	err := addNewUserUseCase.Execute(usecase.AddNewUserUseCaseInput{
		User: usecase.AddUserInput{
			Alias: "dev-user",
			Name:  "Dev User",
			Email: "dev@tig.com",
			Host:  "g",
		},
	})
	if err != nil {
		t.Fatalf("failed creating test user: %v", err)
	}

	checkoutUserUseCase := usecase.NewCheckoutUserUseCase(orgRepo, gitService)
	if err := checkoutUserUseCase.Execute("dev-user"); err != nil {
		t.Fatalf("failed checking out test user: %v", err)
	}

	err = addNewRepoUseCase.Execute(usecase.AddNewRepositoryUseCaseInput{
		Alias:  "my-service",
		Origin: "git@github.com:nirmta/my-service.git",
	})
	if err != nil {
		t.Fatalf("failed adding test repo: %v", err)
	}

	backupUseCase := usecase.NewBackupUseCase(tomlService)
	restoreUseCase := usecase.NewRestoreUseCase(tomlService, orgRepo, gitService)
	backupFile := filepath.Join(tempDir, "test-backup.toml")

	// Act 1: Backup
	backupErr := backupUseCase.Execute(usecase.BackupUseCaseInput{
		DestinationPath: backupFile,
	})

	// Assert 1: Backup file exists and is not empty
	if backupErr != nil {
		t.Fatalf("expected backup to succeed: %v", backupErr)
	}
	info, err := os.Stat(backupFile)
	if err != nil || info.Size() == 0 {
		t.Fatalf("backup file does not exist or is empty")
	}

	// Act 2: Simulate data loss / reset
	removeUserUseCase := usecase.NewRemoveUserUseCase(orgRepo, gitService)
	if err := removeUserUseCase.Execute("dev-user"); err != nil {
		t.Fatalf("failed removing user to simulate data loss: %v", err)
	}
	deletedUser, _ := orgRepo.GetUserByAlias("dev-user")
	if deletedUser != nil {
		t.Fatalf("precondition failed: user should be deleted")
	}

	// Act 3: Restore
	restoreErr := restoreUseCase.Execute(usecase.RestoreUseCaseInput{
		SourcePath: backupFile,
	})

	// Assert 3: User and repo restored
	if restoreErr != nil {
		t.Fatalf("expected restore to succeed: %v", restoreErr)
	}

	restoredUser, err := orgRepo.GetUserByAlias("dev-user")
	if err != nil || restoredUser == nil {
		t.Fatalf("expected user 'dev-user' to be restored")
	}
	if restoredUser.Email() != "dev@tig.com" {
		t.Errorf("expected email 'dev@tig.com', got %s", restoredUser.Email())
	}
	repo := restoredUser.FindRepositoryByAlias("my-service")
	if repo == nil {
		t.Fatalf("expected repository 'my-service' to be restored")
	}
	if repo.Origin() != "dev-user:nirmta/my-service.git" {
		t.Errorf("expected origin 'dev-user:nirmta/my-service.git', got %s", repo.Origin())
	}
}
