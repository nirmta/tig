package cmd

import (
	"tig/app/application/usecase"

	"github.com/urfave/cli/v2"
)

type AddRepositoryCmd struct {
	addNewRepositoryUseCase *usecase.AddNewRepositoryUseCase
}

func NewAddRepositoryCmd(addNewRepositoryUseCase *usecase.AddNewRepositoryUseCase) *AddRepositoryCmd {
	return &AddRepositoryCmd{
		addNewRepositoryUseCase: addNewRepositoryUseCase,
	}
}

func (ar *AddRepositoryCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Flags:   ar.buildFlags(),
		Usage:   "add a repository",
		Action:  ar.buildAction,
	}
}

func (ar *AddRepositoryCmd) buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     FlagAlias,
			Aliases:  []string{FlagAliasShort},
			Usage:    "repository alias name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     FlagOrigin,
			Aliases:  []string{FlagOriginShort},
			Usage:    "repository origin",
			Required: true,
		},
	}
}

func (ar *AddRepositoryCmd) buildAction(cCtx *cli.Context) error {
	return ar.addNewRepositoryUseCase.Execute(usecase.AddNewRepositoryUseCaseInput{
		Alias:  cCtx.String(FlagAlias),
		Origin: cCtx.String(FlagOrigin),
	})
}
