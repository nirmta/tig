package cmd

import (
	"tig/app/application/usecase"

	"github.com/urfave/cli/v2"
)

type CheckoutUserCmd struct {
	checkoutUserUseCase *usecase.CheckoutUserUseCase
}

func NewCheckoutUserCmd(checkoutUserUseCase *usecase.CheckoutUserUseCase) *CheckoutUserCmd {
	return &CheckoutUserCmd{
		checkoutUserUseCase: checkoutUserUseCase,
	}
}

func (co *CheckoutUserCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "checkout",
		Aliases: []string{"co"},
		Flags:   co.buildFlags(),
		Usage:   "checkout user",
		Action:  co.buildAction,
	}
}

func (co *CheckoutUserCmd) buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     FlagAlias,
			Aliases:  []string{FlagAliasShort},
			Usage:    "alias name",
			Required: true,
		},
	}
}

func (co *CheckoutUserCmd) buildAction(cCtx *cli.Context) error {
	userAlias := cCtx.String(FlagAlias)
	return co.checkoutUserUseCase.Execute(userAlias)
}
