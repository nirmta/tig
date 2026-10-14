package settings

import (
	"tig/app/representation/schemas"

	"github.com/urfave/cli/v2"
)

type CmdCollections struct {
	User       []schemas.CmdSchema
	Repository []schemas.CmdSchema
}

func BuildCLI(cmdCollections CmdCollections) []*cli.Command {
	var userCmds []*cli.Command
	for _, userCmd := range cmdCollections.User {
		userCmds = append(userCmds, userCmd.BuildSetting())
	}

	var repositoryCmds []*cli.Command
	for _, repositoryCmd := range cmdCollections.Repository {
		repositoryCmds = append(repositoryCmds, repositoryCmd.BuildSetting())
	}

	return []*cli.Command{
		{
			Name:    "user",
			Aliases: []string{"u"},

			Usage:       "user operations",
			Subcommands: userCmds,
		},
		{
			Name:    "repository",
			Aliases: []string{"r"},

			Usage:       "repository operations",
			Subcommands: repositoryCmds,
		},
	}
}
