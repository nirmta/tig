package main

import (
	"os"
	usecases "tig/app/application/use-cases"
	"tig/app/infrastructure/consts"
	"tig/app/infrastructure/persistence/repositories"
	infrastructureServices "tig/app/infrastructure/services"
	"tig/app/representation/cli/cmds"
	"tig/app/representation/cli/settings"
	"tig/app/representation/schemas"
	representationServices "tig/app/representation/services"

	urfave "github.com/urfave/cli/v2"
)

var jsonService *infrastructureServices.JsonService = infrastructureServices.NewJsonService()
var fsService *infrastructureServices.FsService = infrastructureServices.NewFsService()

func main() {
	preExecution()

	gitService := infrastructureServices.NewGitService(fsService, jsonService)
	organizationRepository := repositories.NewOrganizationRepository(jsonService, fsService, gitService)

	// getRepositoryHostsUseCase := usecases.NewGetRepositoryHostsUseCase()
	newGetUsersUseCase := usecases.NewGetUsersUseCase(organizationRepository, gitService)
	addNewUserUseCase := usecases.NewAddNewUserUseCase(organizationRepository, gitService)
	removeUserUseCase := usecases.NewRemoveUserUseCase(organizationRepository, gitService)
	checkoutUserUseCase := usecases.NewCheckoutUserUseCase(organizationRepository, gitService)
	getUserGitConfigUseCase := usecases.NewGetUserGitConfigUseCase(organizationRepository, gitService)
	generateKeysUserUseCase := usecases.NewGenerateKeysUserUseCase(organizationRepository, gitService)

	addRepositoryUseCase := usecases.NewAddNewRepositoryUseCase(organizationRepository, gitService)
	removeRepoistoryUseCase := usecases.NewRemoveRepositoryUseCase(organizationRepository, gitService)
	getRepositoriesUseCase := usecases.NewGetRepositoriesUseCase(organizationRepository)
	cloneRepositoryUseCase := usecases.NewDownloadRepositoriesUseCase(organizationRepository, gitService, fsService)

	printService := representationServices.NewPrintService()

	var addUserCmd schemas.CmdSchema = cmds.NewAddUserCmd(addNewUserUseCase)
	var removeUserCmd schemas.CmdSchema = cmds.NewRemoveUserCmd(removeUserUseCase)
	var listUserCmd schemas.CmdSchema = cmds.NewListUsersCmd(newGetUsersUseCase, printService)
	var checkoutUserCmd schemas.CmdSchema = cmds.NewCheckoutUserCmd(checkoutUserUseCase)
	var listUserGitConfigCmd schemas.CmdSchema = cmds.NewListUserGitConfigCmd(getUserGitConfigUseCase, printService)

	var addRepositoryCmd schemas.CmdSchema = cmds.NewAddRepositoryCmd(addRepositoryUseCase)
	var removeRepositoryCmd schemas.CmdSchema = cmds.NewRemoveRepositoryCmd(removeRepoistoryUseCase)
	var listRepositoriesCmd schemas.CmdSchema = cmds.NewListRepositoriesCmd(getRepositoriesUseCase, printService)
	var cloneRepositoriesCmd schemas.CmdSchema = cmds.NewCloneRepositoriesCmd(cloneRepositoryUseCase)
	var generateKeysUser schemas.CmdSchema = cmds.NewGenerateKeysUser(generateKeysUserUseCase)

	commands := settings.BuildCLI(settings.CmdCollections{
		User:       []schemas.CmdSchema{addUserCmd, removeUserCmd, listUserCmd, checkoutUserCmd, listUserGitConfigCmd, generateKeysUser},
		Repository: []schemas.CmdSchema{addRepositoryCmd, removeRepositoryCmd, listRepositoriesCmd, cloneRepositoriesCmd},
	})

	app := &urfave.App{
		Commands: commands,
		Version:  os.Getenv("VERSION"),
	}

	app.Run(os.Args)
}

func preExecution() {
	systemFolder := fsService.GetFolderPath([]string{consts.FOLDER})
	fsService.CreateDirectory(systemFolder)

	jsonPath := fsService.GetFolderPath([]string{consts.FOLDER, jsonService.AddExtension(consts.FILE_NAME)})

	if existedFile := fsService.ExistFile(jsonPath); existedFile {
		return
	}

	jsonService.CreateFile("", systemFolder, consts.FILE_NAME)
}
