package cmd

import (
	"tig/app/application/usecase"

	"github.com/urfave/cli/v2"
)

type RemoveRepositoryCmd struct {
	removeRepositoryUseCase *usecase.RemoveRepositoryUseCase
}

func NewRemoveRepositoryCmd(removeRepositoryUseCase *usecase.RemoveRepositoryUseCase) *RemoveRepositoryCmd {
	return &RemoveRepositoryCmd{
		removeRepositoryUseCase: removeRepositoryUseCase,
	}
}

func (rr *RemoveRepositoryCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"rm"},
		Flags:   rr.buildFlags(),
		Usage:   "remove a repository",
		Action:  rr.buildAction,
	}
}

func (rr *RemoveRepositoryCmd) buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     FlagAlias,
			Aliases:  []string{FlagAliasShort},
			Usage:    "alias repository",
			Required: true,
		},
	}
}

func (rr *RemoveRepositoryCmd) buildAction(cCtx *cli.Context) error {
	repositoryAlias := cCtx.String(FlagAlias)
	return rr.removeRepositoryUseCase.Execute(repositoryAlias)
}
