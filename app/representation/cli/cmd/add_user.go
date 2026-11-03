package cmd

import (
	"tig/app/application/usecase"

	"github.com/urfave/cli/v2"
)

type AddUserCmd struct {
	addNewUserUseCase *usecase.AddNewUserUseCase
}

func NewAddUserCmd(addNewUserUseCase *usecase.AddNewUserUseCase) *AddUserCmd {
	return &AddUserCmd{
		addNewUserUseCase: addNewUserUseCase,
	}
}

func (au *AddUserCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Flags:   au.buildFlags(),
		Usage:   "add a user",
		Action:  au.buildAction,
	}
}

func (au *AddUserCmd) buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     FlagAlias,
			Aliases:  []string{FlagAliasShort},
			Usage:    "user alias name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     FlagEmail,
			Aliases:  []string{FlagEmailShort},
			Usage:    "user email",
			Required: true,
		},
		&cli.StringFlag{
			Name:     FlagName,
			Aliases:  []string{FlagNameShort},
			Usage:    "user name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     FlagHost,
			Aliases:  []string{FlagHostShort},
			Usage:    "repository host",
			Required: true,
		},
	}
}

func (au *AddUserCmd) buildAction(cCtx *cli.Context) error {
	addNewUserUseCaseInput := usecase.AddNewUserUseCaseInput{
		User: usecase.AddUserInput{
			Alias: cCtx.String(FlagAlias),
			Name:  cCtx.String(FlagName),
			Email: cCtx.String(FlagEmail),
			Host:  cCtx.String(FlagHost),
		},
	}

	return au.addNewUserUseCase.Execute(addNewUserUseCaseInput)
}
