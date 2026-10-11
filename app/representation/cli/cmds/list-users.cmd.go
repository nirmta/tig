package cmds

import (
	usecases "tig/app/application/use-cases"
	"tig/app/representation/services"

	"github.com/urfave/cli/v2"
)

type ListUsersCmd struct {
	getUsersUseCase *usecases.GetUsersUseCase
	printService    *services.PrintService
}

func NewListUsersCmd(
	getUsersUseCase *usecases.GetUsersUseCase,
	printService *services.PrintService,
) *ListUsersCmd {
	return &ListUsersCmd{
		getUsersUseCase: getUsersUseCase,
		printService:    printService,
	}
}

func (ld *ListUsersCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Flags:   ld.BuildFlags(),
		Usage:   "list users",
		Action:  ld.BuildAction,
	}
}

func (ld *ListUsersCmd) BuildFlags() []cli.Flag {
	return []cli.Flag{}
}

func (ld *ListUsersCmd) BuildAction(cCtx *cli.Context) error {
	users := ld.getUsersUseCase.Execute()

	table := [][]string{
		{"Alias", "User Name", "Email", "Repository"},
	}

	for _, user := range users {
		name := user.Name()

		if isGlobal := user.InGlobal(); isGlobal {
			name = "*" + name
		}

		table = append(table, []string{user.Alias(), name, user.Email(), user.RepositoryHost()})
	}

	ld.printService.RenderTable(table)

	return nil
}
