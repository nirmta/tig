//go:build wireinject
// +build wireinject

package bootstrap

import (
	"fmt"
	"tig/app/application/port"
	"tig/app/application/repository"
	"tig/app/application/usecase"
	"tig/app/domain/entities/organization"
	"tig/app/infrastructure/consts"
	"tig/app/infrastructure/persistence/repositories"
	"tig/app/infrastructure/services"
	"tig/app/representation/cli"
	"tig/app/representation/presenter"

	"github.com/goforj/wire"
	urfave "github.com/urfave/cli/v2"
)

var Version = "dev"

type Workspace struct {
	FsService   *services.FsService
	TomlService *services.TomlService
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

func provideWorkspace() (*Workspace, error) {
	fsService := services.NewFsService()
	tomlService := services.NewTomlService()

	if err := initWorkspace(fsService, tomlService); err != nil {
		return nil, fmt.Errorf("failed to init workspace: %w", err)
	}

	return &Workspace{
		FsService:   fsService,
		TomlService: tomlService,
	}, nil
}

var workspaceSet = wire.NewSet(
	provideWorkspace,
	wire.FieldsOf(new(*Workspace), "FsService", "TomlService"),
	wire.Bind(new(port.FsService), new(*services.FsService)),
	wire.Bind(new(port.TomlService), new(*services.TomlService)),
)

var serviceSet = wire.NewSet(
	workspaceSet,
	services.NewGitService,
	wire.Bind(new(port.GitService), new(*services.GitService)),
	presenter.NewPrintService,
)

var repositorySet = wire.NewSet(
	repositories.NewOrganizationRepository,
	wire.Bind(new(repository.OrganizationRepository), new(*repositories.OrganizationRepository)),
)

var useCaseSet = wire.NewSet(
	usecase.NewGetUsersUseCase,
	usecase.NewAddNewUserUseCase,
	usecase.NewRemoveUserUseCase,
	usecase.NewCheckoutUserUseCase,
	usecase.NewGetUserGitConfigUseCase,
	usecase.NewGenerateKeysUserUseCase,
	usecase.NewAddNewRepositoryUseCase,
	usecase.NewRemoveRepositoryUseCase,
	usecase.NewGetRepositoriesUseCase,
	usecase.NewDownloadRepositoriesUseCase,
	usecase.NewBackupUseCase,
	usecase.NewRestoreUseCase,
)

var cmdSet = wire.NewSet(
	cli.NewAddUserCmd,
	cli.NewRemoveUserCmd,
	cli.NewListUsersCmd,
	cli.NewCheckoutUserCmd,
	cli.NewListUserGitConfigCmd,
	cli.NewGenerateKeysUser,
	cli.NewAddRepositoryCmd,
	cli.NewRemoveRepositoryCmd,
	cli.NewListRepositoriesCmd,
	cli.NewCloneRepositoriesCmd,
	cli.NewBackupCmd,
	cli.NewRestoreCmd,
	provideCommandsList,
	cli.Build,
)

func provideCommandsList(
	addUserCmd *cli.AddUserCmd,
	removeUserCmd *cli.RemoveUserCmd,
	listUserCmd *cli.ListUsersCmd,
	checkoutUserCmd *cli.CheckoutUserCmd,
	listUserGitConfigCmd *cli.ListUserGitConfigCmd,
	generateKeysUser *cli.GenerateKeysUser,
	addRepositoryCmd *cli.AddRepositoryCmd,
	removeRepositoryCmd *cli.RemoveRepositoryCmd,
	listRepositoriesCmd *cli.ListRepositoriesCmd,
	cloneRepositoriesCmd *cli.CloneRepositoriesCmd,
	backupCmd *cli.BackupCmd,
	restoreCmd *cli.RestoreCmd,
) []cli.Command {
	return []cli.Command{
		addUserCmd,
		removeUserCmd,
		listUserCmd,
		checkoutUserCmd,
		listUserGitConfigCmd,
		generateKeysUser,
		addRepositoryCmd,
		removeRepositoryCmd,
		listRepositoriesCmd,
		cloneRepositoriesCmd,
		backupCmd,
		restoreCmd,
	}
}

func provideApp(commands []*urfave.Command) *urfave.App {
	return &urfave.App{
		Name:        "tig",
		Usage:       "Git profile, SSH key, and repository manager",
		Description: "A developer CLI tool to seamlessly switch between multiple Git accounts, manage SSH keys, and clone repositories.",
		Version:     Version,
		Commands:    commands,
	}
}

func InitializeApp() (*urfave.App, error) {
	wire.Build(
		serviceSet,
		repositorySet,
		useCaseSet,
		cmdSet,
		provideApp,
	)
	return nil, nil
}
