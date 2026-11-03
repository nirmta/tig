package cmd

import (
	"tig/app/application/usecase"

	"github.com/urfave/cli/v2"
)

type CloneRepositoriesCmd struct {
	downloadRepositoriesUseCase *usecase.DownloadRepositoriesUseCase
}

func NewCloneRepositoriesCmd(addNewRepositoryUseCase *usecase.DownloadRepositoriesUseCase) *CloneRepositoriesCmd {
	return &CloneRepositoriesCmd{
		downloadRepositoriesUseCase: addNewRepositoryUseCase,
	}
}

func (ar *CloneRepositoriesCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "clone",
		Aliases: []string{"c"},
		Flags:   ar.BuildFlags(),
		Usage:   "clone the user repositories",
		Action:  ar.BuildAction,
	}
}

func (ar *CloneRepositoriesCmd) BuildFlags() []cli.Flag {
	return []cli.Flag{}
}

func (ar *CloneRepositoriesCmd) BuildAction(cCtx *cli.Context) error {
	return ar.downloadRepositoriesUseCase.Execute()
}
