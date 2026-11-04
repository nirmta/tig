package cli

import "github.com/urfave/cli/v2"

type Command interface {
	BuildSetting() *cli.Command
}

type CmdCollections struct {
	User       []Command
	Repository []Command
	Root       []Command
}

func Build(cmdCollections CmdCollections) []*cli.Command {
	var userCmds []*cli.Command
	for _, userCmd := range cmdCollections.User {
		userCmds = append(userCmds, userCmd.BuildSetting())
	}

	var repositoryCmds []*cli.Command
	for _, repositoryCmd := range cmdCollections.Repository {
		repositoryCmds = append(repositoryCmds, repositoryCmd.BuildSetting())
	}

	commands := []*cli.Command{
		{
			Name:        "user",
			Aliases:     []string{"u"},
			Usage:       "user operations",
			Subcommands: userCmds,
		},
		{
			Name:        "repository",
			Aliases:     []string{"r"},
			Usage:       "repository operations",
			Subcommands: repositoryCmds,
		},
	}

	for _, root := range cmdCollections.Root {
		commands = append(commands, root.BuildSetting())
	}

	return commands
}
