package main

import (
	"fmt"
	"os"
	"tig/app/application/usecase"
	"tig/app/domain/entities/organization"
	"tig/app/infrastructure/consts"
	"tig/app/infrastructure/persistence/repositories"
	"tig/app/infrastructure/services"
	"tig/app/representation/cli"
	"tig/app/representation/cli/cmd"
	"tig/app/representation/presenter"

	urfave "github.com/urfave/cli/v2"
)

var Version = "dev"

func main() {
	tomlService := services.NewTomlService()
	fsService := services.NewFsService()

	if err := initWorkspace(fsService, tomlService); err != nil {
		fmt.Fprintf(os.Stderr, "Initialization error: %v\n", err)
		os.Exit(1)
	}

	gitService := services.NewGitService(fsService)
	organizationRepository := repositories.NewOrganizationRepository(tomlService, fsService, gitService)

	newGetUsersUseCase := usecase.NewGetUsersUseCase(organizationRepository, gitService)
	addNewUserUseCase := usecase.NewAddNewUserUseCase(organizationRepository, gitService)
	removeUserUseCase := usecase.NewRemoveUserUseCase(organizationRepository, gitService)
	checkoutUserUseCase := usecase.NewCheckoutUserUseCase(organizationRepository, gitService)
	getUserGitConfigUseCase := usecase.NewGetUserGitConfigUseCase(organizationRepository, gitService)
	generateKeysUserUseCase := usecase.NewGenerateKeysUserUseCase(organizationRepository, gitService)

	addRepositoryUseCase := usecase.NewAddNewRepositoryUseCase(organizationRepository, gitService)
	removeRepositoryUseCase := usecase.NewRemoveRepositoryUseCase(organizationRepository, gitService)
	getRepositoriesUseCase := usecase.NewGetRepositoriesUseCase(organizationRepository)
	cloneRepositoryUseCase := usecase.NewDownloadRepositoriesUseCase(organizationRepository, gitService, fsService)
	backupUseCase := usecase.NewBackupUseCase(tomlService)
	restoreUseCase := usecase.NewRestoreUseCase(tomlService, organizationRepository, gitService)

	printService := presenter.NewPrintService()

	addUserCmd := cmd.NewAddUserCmd(addNewUserUseCase)
	removeUserCmd := cmd.NewRemoveUserCmd(removeUserUseCase)
	listUserCmd := cmd.NewListUsersCmd(newGetUsersUseCase, printService)
	checkoutUserCmd := cmd.NewCheckoutUserCmd(checkoutUserUseCase)
	listUserGitConfigCmd := cmd.NewListUserGitConfigCmd(getUserGitConfigUseCase, printService)
	generateKeysUser := cmd.NewGenerateKeysUser(generateKeysUserUseCase)

	addRepositoryCmd := cmd.NewAddRepositoryCmd(addRepositoryUseCase)
	removeRepositoryCmd := cmd.NewRemoveRepositoryCmd(removeRepositoryUseCase)
	listRepositoriesCmd := cmd.NewListRepositoriesCmd(getRepositoriesUseCase, printService)
	cloneRepositoriesCmd := cmd.NewCloneRepositoriesCmd(cloneRepositoryUseCase)

	backupCmd := cmd.NewBackupCmd(backupUseCase, printService)
	restoreCmd := cmd.NewRestoreCmd(restoreUseCase, printService)

	commands := cli.Build(cli.CmdCollections{
		User:       []cli.Command{addUserCmd, removeUserCmd, listUserCmd, checkoutUserCmd, listUserGitConfigCmd, generateKeysUser},
		Repository: []cli.Command{addRepositoryCmd, removeRepositoryCmd, listRepositoriesCmd, cloneRepositoriesCmd},
		Root:       []cli.Command{backupCmd, restoreCmd},
	})

	app := &urfave.App{
		Name:        "tig",
		Usage:       "Git profile, SSH key, and repository manager",
		Description: "A developer CLI tool to seamlessly switch between multiple Git accounts, manage SSH keys, and clone repositories.",
		Version:     Version,
		Commands:    commands,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func initWorkspace(fsService *services.FsService, tomlService *services.TomlService) error {
	systemFolder := fsService.GetFolderPath([]string{consts.FOLDER})
	if err := fsService.CreateDirectory(systemFolder); err != nil {
		return fmt.Errorf("failed to create config directory %q: %w", systemFolder, err)
	}

	tomlPath := fsService.GetFolderPath([]string{consts.FOLDER, tomlService.AddExtension(consts.FILE_NAME)})
	if !fsService.ExistFile(tomlPath) {
		if err := tomlService.CreateFile(organization.New(), systemFolder, consts.FILE_NAME); err != nil {
			return fmt.Errorf("failed to create config file %q: %w", tomlPath, err)
		}
	}

	return nil
}
