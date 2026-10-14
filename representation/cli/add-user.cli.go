package cli

import (
	usecases "tig/application/use-cases"
	"tig/representation/mappers"

	"github.com/urfave/cli/v2"
)

type AddUserCli struct {
	addNewUserUseCase *usecases.AddNewUserUseCase
}

func NewAddUserCli(addNewUserUseCase *usecases.AddNewUserUseCase) *AddUserCli {
	return &AddUserCli{
		addNewUserUseCase: addNewUserUseCase,
	}
}

func (au *AddUserCli) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Flags:   au.BuildFlags(),
		Usage:   "add a user",
		Action:  au.BuildAction,
	}
}

func (au *AddUserCli) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     mappers.Alias,
			Aliases:  []string{"a"},
			Usage:    "alias name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     mappers.Email,
			Aliases:  []string{"e"},
			Usage:    "email",
			Required: true,
		},
		&cli.StringFlag{
			Name:     mappers.Name,
			Aliases:  []string{"n"},
			Usage:    "user name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     mappers.Host,
			Aliases:  []string{"r"},
			Usage:    "repository host",
			Required: true,
		},
	}
}

func (au *AddUserCli) BuildAction(cCtx *cli.Context) error {
	addNewUserUseCaseInput := usecases.AddNewUserUseCaseInput{
		User: usecases.AddUserInput{
			Alias: cCtx.String(mappers.Alias),
			Name:  cCtx.String(mappers.Name),
			Email: cCtx.String(mappers.Email),
			Host:  cCtx.String(mappers.Host),
		},
	}

	au.addNewUserUseCase.Execute(addNewUserUseCaseInput)
	return nil
}
