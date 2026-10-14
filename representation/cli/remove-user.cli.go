package cli

import (
	usecases "tig/application/use-cases"
	"tig/representation/mappers"

	"github.com/urfave/cli/v2"
)

type RemoveUserCli struct {
	removeUserUseCase *usecases.RemoveUserUseCase
}

func NewRemoveUserCli(removeUserUseCase *usecases.RemoveUserUseCase) *RemoveUserCli {
	return &RemoveUserCli{
		removeUserUseCase: removeUserUseCase,
	}
}

func (ru *RemoveUserCli) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"r"},
		Flags:   ru.BuildFlags(),
		Usage:   "remove a user",
		Action:  ru.BuildAction,
	}
}

func (ru *RemoveUserCli) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     mappers.Alias,
			Aliases:  []string{"a"},
			Usage:    "alias name",
			Required: true,
		},
	}
}

func (ru *RemoveUserCli) BuildAction(cCtx *cli.Context) error {
	userAlias := cCtx.String(mappers.Alias)
	ru.removeUserUseCase.Execute(userAlias)
	return nil
}
