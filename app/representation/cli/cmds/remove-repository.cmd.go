package cmds

import (
	usecases "tig/app/application/use-cases"
	"tig/app/representation/cli/aliases"
	"tig/app/representation/cli/flags"

	"github.com/urfave/cli/v2"
)

type RemoveRepositoryCmd struct {
	removeRepositoryUseCase *usecases.RemoveRepositoryUseCase
}

func NewRemoveRepositoryCmd(removeRepositoryUseCase *usecases.RemoveRepositoryUseCase) *RemoveRepositoryCmd {
	return &RemoveRepositoryCmd{
		removeRepositoryUseCase: removeRepositoryUseCase,
	}
}

func (rr *RemoveRepositoryCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"r"},
		Flags:   rr.BuildFlags(),
		Usage:   "remove a repository",
		Action:  rr.BuildAction,
	}
}

func (rr *RemoveRepositoryCmd) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     flags.RepositoryAlias,
			Aliases:  []string{aliases.RepositoryAlias},
			Usage:    "alias repository",
			Required: true,
		},
	}
}

func (rr *RemoveRepositoryCmd) BuildAction(cCtx *cli.Context) error {
	repositoryAlias := cCtx.String(flags.RepositoryAlias)
	_, err := rr.removeRepositoryUseCase.Execute(repositoryAlias)
	return err
}
