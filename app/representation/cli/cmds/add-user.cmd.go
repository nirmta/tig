package cmds

import (
	usecases "tig/app/application/use-cases"
	"tig/app/representation/cli/aliases"
	"tig/app/representation/cli/flags"

	"github.com/urfave/cli/v2"
)

type AddUserCmd struct {
	addNewUserUseCase *usecases.AddNewUserUseCase
}

func NewAddUserCmd(addNewUserUseCase *usecases.AddNewUserUseCase) *AddUserCmd {
	return &AddUserCmd{
		addNewUserUseCase: addNewUserUseCase,
	}
}

func (au *AddUserCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Flags:   au.BuildFlags(),
		Usage:   "add a user",
		Action:  au.BuildAction,
	}
}

func (au *AddUserCmd) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     flags.UserAlias,
			Aliases:  []string{aliases.UserAlias},
			Usage:    "user alias name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     flags.UserEmail,
			Aliases:  []string{aliases.UserEmail},
			Usage:    "user email",
			Required: true,
		},
		&cli.StringFlag{
			Name:     flags.UserName,
			Aliases:  []string{aliases.UserName},
			Usage:    "user name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     flags.UserHost,
			Aliases:  []string{aliases.UserHost},
			Usage:    "repository host",
			Required: true,
		},
	}
}

func (au *AddUserCmd) BuildAction(cCtx *cli.Context) error {
	addNewUserUseCaseInput := usecases.AddNewUserUseCaseInput{
		User: usecases.AddUserInput{
			Alias: cCtx.String(flags.UserAlias),
			Name:  cCtx.String(flags.UserName),
			Email: cCtx.String(flags.UserEmail),
			Host:  cCtx.String(flags.UserHost),
		},
	}

	_, err := au.addNewUserUseCase.Execute(addNewUserUseCaseInput)
	return err
}
