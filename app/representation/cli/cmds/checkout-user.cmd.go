package cmds

import (
	usecases "tig/app/application/use-cases"
	"tig/app/representation/cli/flags"

	"github.com/urfave/cli/v2"
)

type CheckoutUserCmd struct {
	checkoutUserUseCase *usecases.CheckoutUserUseCase
}

func NewCheckoutUserCmd(checkoutUserUseCase *usecases.CheckoutUserUseCase) *CheckoutUserCmd {
	return &CheckoutUserCmd{
		checkoutUserUseCase: checkoutUserUseCase,
	}
}

func (co *CheckoutUserCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "checkout",
		Aliases: []string{"c"},
		Flags:   co.BuildFlags(),
		Usage:   "checkout user",
		Action:  co.BuildAction,
	}
}

func (co *CheckoutUserCmd) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     flags.UserAlias,
			Aliases:  []string{"a"},
			Usage:    "alias name",
			Required: true,
		},
	}
}

func (co *CheckoutUserCmd) BuildAction(cCtx *cli.Context) error {
	userAlias := cCtx.String(flags.UserAlias)
	co.checkoutUserUseCase.Execute(userAlias)
	return nil
}
