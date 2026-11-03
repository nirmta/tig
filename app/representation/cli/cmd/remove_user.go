package cmd

import (
	"tig/app/application/usecase"

	"github.com/urfave/cli/v2"
)

type RemoveUserCmd struct {
	removeUserUseCase *usecase.RemoveUserUseCase
}

func NewRemoveUserCmd(removeUserUseCase *usecase.RemoveUserUseCase) *RemoveUserCmd {
	return &RemoveUserCmd{
		removeUserUseCase: removeUserUseCase,
	}
}

func (ru *RemoveUserCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"rm"},
		Flags:   ru.buildFlags(),
		Usage:   "remove a user",
		Action:  ru.buildAction,
	}
}

func (ru *RemoveUserCmd) buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     FlagAlias,
			Aliases:  []string{FlagAliasShort},
			Usage:    "alias name",
			Required: true,
		},
	}
}

func (ru *RemoveUserCmd) buildAction(cCtx *cli.Context) error {
	userAlias := cCtx.String(FlagAlias)
	return ru.removeUserUseCase.Execute(userAlias)
}
