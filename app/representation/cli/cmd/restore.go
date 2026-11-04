package cmd

import (
	"fmt"
	"tig/app/application/usecase"
	"tig/app/representation/presenter"

	"github.com/urfave/cli/v2"
)

type RestoreCmd struct {
	restoreUseCase *usecase.RestoreUseCase
	printService   *presenter.PrintService
}

func NewRestoreCmd(restoreUseCase *usecase.RestoreUseCase, printService *presenter.PrintService) *RestoreCmd {
	return &RestoreCmd{
		restoreUseCase: restoreUseCase,
		printService:   printService,
	}
}

func (rc *RestoreCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "restore",
		Flags:   rc.buildFlags(),
		Usage:   "restore tig configuration from a backup toml file",
		Action:  rc.buildAction,
	}
}

func (rc *RestoreCmd) buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    FlagInput,
			Aliases: []string{FlagInputShort},
			Usage:   "path to the backup toml file",
		},
	}
}

func (rc *RestoreCmd) buildAction(cCtx *cli.Context) error {
	srcPath := cCtx.String(FlagInput)
	if srcPath == "" && cCtx.Args().Present() {
		srcPath = cCtx.Args().First()
	}

	if srcPath == "" {
		return fmt.Errorf("source backup file path is required")
	}

	err := rc.restoreUseCase.Execute(usecase.RestoreUseCaseInput{
		SourcePath: srcPath,
	})
	if err != nil {
		return err
	}

	rc.printService.Println("✔ Configuration restored successfully and ~/.ssh/config regenerated.")
	return nil
}
