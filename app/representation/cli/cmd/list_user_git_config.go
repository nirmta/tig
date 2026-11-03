package cmd

import (
	"tig/app/application/usecase"
	"tig/app/representation/presenter"

	"github.com/urfave/cli/v2"
)

type ListUserGitConfigCmd struct {
	getUserGitConfigUseCase *usecase.GetUserGitConfigUseCase
	printService            *presenter.PrintService
}

func NewListUserGitConfigCmd(
	getUserGitConfigUseCase *usecase.GetUserGitConfigUseCase,
	printService *presenter.PrintService,
) *ListUserGitConfigCmd {
	return &ListUserGitConfigCmd{
		getUserGitConfigUseCase: getUserGitConfigUseCase,
		printService:            printService,
	}
}

func (lugc *ListUserGitConfigCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "config",
		Aliases: []string{"s"},
		Flags:   lugc.BuildFlags(),
		Usage:   "list user git configs",
		Action:  lugc.BuildAction,
	}
}

func (lugc *ListUserGitConfigCmd) BuildFlags() []cli.Flag {
	return []cli.Flag{}
}

func (lugc *ListUserGitConfigCmd) BuildAction(cCtx *cli.Context) error {
	userGitConfig, err := lugc.getUserGitConfigUseCase.Execute()

	if err != nil {
		return err
	}

	lugc.printService.Println("=== PUBLIC KEY ===")
	lugc.printService.Println(userGitConfig.SSHConfig.PublicKey)
	lugc.printService.Println("===    END     ===")

	return nil
}
