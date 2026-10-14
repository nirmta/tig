package cmds

import (
	usecases "tig/app/application/use-cases"
	"tig/app/representation/cli/aliases"
	"tig/app/representation/cli/flags"

	"github.com/urfave/cli/v2"
)

type RemoveUserCmd struct {
	removeUserUseCase *usecases.RemoveUserUseCase
}

func NewRemoveUserCmd(removeUserUseCase *usecases.RemoveUserUseCase) *RemoveUserCmd {
	return &RemoveUserCmd{
		removeUserUseCase: removeUserUseCase,
	}
}

func (ru *RemoveUserCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"r"},
		Flags:   ru.BuildFlags(),
		Usage:   "remove a user",
		Action:  ru.BuildAction,
	}
}

func (ru *RemoveUserCmd) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     flags.UserAlias,
			Aliases:  []string{aliases.UserAlias},
			Usage:    "alias name",
			Required: true,
		},
	}
}

func (ru *RemoveUserCmd) BuildAction(cCtx *cli.Context) error {
	userAlias := cCtx.String(flags.UserAlias)
	_, err := ru.removeUserUseCase.Execute(userAlias)
	return err
}
