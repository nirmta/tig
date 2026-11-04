package cmd

import (
	"fmt"
	"tig/app/application/usecase"
	"tig/app/representation/presenter"

	"github.com/urfave/cli/v2"
)

type BackupCmd struct {
	backupUseCase *usecase.BackupUseCase
	printService  *presenter.PrintService
}

func NewBackupCmd(backupUseCase *usecase.BackupUseCase, printService *presenter.PrintService) *BackupCmd {
	return &BackupCmd{
		backupUseCase: backupUseCase,
		printService:  printService,
	}
}

func (bc *BackupCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "backup",
		Aliases: []string{"b"},
		Flags:   bc.buildFlags(),
		Usage:   "backup tig configuration to a toml file",
		Action:  bc.buildAction,
	}
}

func (bc *BackupCmd) buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    FlagOutput,
			Aliases: []string{FlagOutputShort},
			Usage:   "destination file path (default: organization-backup.toml)",
		},
	}
}

func (bc *BackupCmd) buildAction(cCtx *cli.Context) error {
	destPath := cCtx.String(FlagOutput)
	if destPath == "" && cCtx.Args().Present() {
		destPath = cCtx.Args().First()
	}

	err := bc.backupUseCase.Execute(usecase.BackupUseCaseInput{
		DestinationPath: destPath,
	})
	if err != nil {
		return err
	}

	if destPath == "" {
		destPath = "organization-backup.toml"
	}
	bc.printService.Println(fmt.Sprintf("✔ Backup created successfully at %s", destPath))
	return nil
}
