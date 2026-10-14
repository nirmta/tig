package cmds

import (
	usecases "tig/app/application/use-cases"

	"github.com/urfave/cli/v2"
)

type GenerateKeysUser struct {
	generateKeysUserUseCase *usecases.GenerateKeysUserUseCase
}

func NewGenerateKeysUser(generateKeysUserUseCase *usecases.GenerateKeysUserUseCase) *GenerateKeysUser {
	return &GenerateKeysUser{
		generateKeysUserUseCase: generateKeysUserUseCase,
	}
}

func (gku *GenerateKeysUser) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "generate",
		Aliases: []string{"g"},
		Flags:   gku.BuildFlags(),
		Usage:   "generate user ssh keys",
		Action:  gku.BuildAction,
	}
}

func (gku *GenerateKeysUser) BuildFlags() []cli.Flag {
	return []cli.Flag{}
}

func (gku *GenerateKeysUser) BuildAction(cCtx *cli.Context) error {
	_, err := gku.generateKeysUserUseCase.Execute()
	return err
}
