package schemas

import "github.com/urfave/cli/v2"

type CmdSchema interface {
	BuildSetting() *cli.Command
	BuildFlags() []cli.Flag
	BuildAction(cCtx *cli.Context) error
}
