package cmd

import (
	"tig/app/application/usecase"
	"tig/app/representation/presenter"

	"github.com/urfave/cli/v2"
)

type ListRepositoriesCmd struct {
	getRepositoriesUseCase *usecase.GetRepositoriesUseCase
	printService           *presenter.PrintService
}

func NewListRepositoriesCmd(
	getRepositoriesUseCase *usecase.GetRepositoriesUseCase,
	printService *presenter.PrintService,
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
