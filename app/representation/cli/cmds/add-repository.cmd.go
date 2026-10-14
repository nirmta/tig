package cmds

import (
	usecases "tig/app/application/use-cases"
	"tig/app/representation/cli/aliases"
	"tig/app/representation/cli/flags"

	"github.com/urfave/cli/v2"
)

type AddRepositoryCmd struct {
	addNewRepositoryUseCase *usecases.AddNewRepositoryUseCase
}

func NewAddRepositoryCmd(addNewRepositoryUseCase *usecases.AddNewRepositoryUseCase) *AddRepositoryCmd {
	return &AddRepositoryCmd{
		addNewRepositoryUseCase: addNewRepositoryUseCase,
	}
}

func (ar *AddRepositoryCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Flags:   ar.BuildFlags(),
		Usage:   "add a repository",
		Action:  ar.BuildAction,
	}
}

func (ar *AddRepositoryCmd) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     flags.RepositoryAlias,
			Aliases:  []string{aliases.RepositoryAlias},
			Usage:    "repository alias name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     flags.RepositoryOrigin,
			Aliases:  []string{aliases.RepositoryOrigin},
			Usage:    "repository origin",
			Required: true,
		},
	}
}

func (ar *AddRepositoryCmd) BuildAction(cCtx *cli.Context) error {
	_, err := ar.addNewRepositoryUseCase.Execute(usecases.AddNewRepositoryUseCaseInput{
		Alias:  cCtx.String(flags.RepositoryAlias),
		Origin: cCtx.String(flags.RepositoryOrigin),
	})

	return err
}
