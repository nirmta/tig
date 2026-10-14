package cmds

import (
	usecases "tig/app/application/use-cases"
	"tig/app/representation/services"

	"github.com/urfave/cli/v2"
)

type ListRepositoriesCmd struct {
	getRepositoriesUseCase *usecases.GetRepositoriesUseCase
	printService           *services.PrintService
}

func NewListRepositoriesCmd(
	getRepositoriesUseCase *usecases.GetRepositoriesUseCase,
	printService *services.PrintService,
) *ListRepositoriesCmd {
	return &ListRepositoriesCmd{
		getRepositoriesUseCase: getRepositoriesUseCase,
		printService:           printService,
	}
}

func (ld *ListRepositoriesCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Flags:   ld.BuildFlags(),
		Usage:   "list repositories",
		Action:  ld.BuildAction,
	}
}

func (ld *ListRepositoriesCmd) BuildFlags() []cli.Flag {
	return []cli.Flag{}
}

func (ld *ListRepositoriesCmd) BuildAction(cCtx *cli.Context) error {
	repositories, err := ld.getRepositoriesUseCase.Execute()

	if err != nil {
		return err
	}

	table := [][]string{
		{"Alias", "Origin"},
	}

	for _, repository := range repositories {
		table = append(table, []string{repository.Alias(), repository.Origin()})
	}

	ld.printService.RenderTable(table)
	return nil
}
