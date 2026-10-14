package cli

import (
	usecases "tig/application/use-cases"
	"tig/representation/mappers"

	"github.com/urfave/cli/v2"
)

type CheckoutUserCli struct {
	checkoutUserUseCase *usecases.CheckoutUserUseCase
}

func NewCheckoutUserCli(checkoutUserUseCase *usecases.CheckoutUserUseCase) *CheckoutUserCli {
	return &CheckoutUserCli{
		checkoutUserUseCase: checkoutUserUseCase,
	}
}

func (co *CheckoutUserCli) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "checkout",
		Aliases: []string{"c"},
		Flags:   co.BuildFlags(),
		Usage:   "checkout user",
		Action:  co.BuildAction,
	}
}

func (co *CheckoutUserCli) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     mappers.Alias,
			Aliases:  []string{"a"},
			Usage:    "alias name",
			Required: true,
		},
	}
}

func (co *CheckoutUserCli) BuildAction(cCtx *cli.Context) error {
	userAlias := cCtx.String(mappers.Alias)
	co.checkoutUserUseCase.Execute(userAlias)
	return nil
}
