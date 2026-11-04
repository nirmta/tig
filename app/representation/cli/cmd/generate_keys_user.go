package cmd

import (
	"tig/app/application/usecase"

	"github.com/urfave/cli/v2"
)

type GenerateKeysUser struct {
	generateKeysUserUseCase *usecase.GenerateKeysUserUseCase
}

func NewGenerateKeysUser(generateKeysUserUseCase *usecase.GenerateKeysUserUseCase) *GenerateKeysUser {
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
	return gku.generateKeysUserUseCase.Execute()
}
