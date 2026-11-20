package cli

import "github.com/urfave/cli/v2"

type Command interface {
	Build() *cli.Command
}

func Build(commands []Command) []*cli.Command {
	var commandsByCategory = make(map[Category][]*cli.Command)

	for _, command := range commands {
		builtCommand := command.Build()
		category := Category(builtCommand.Category)
		commandsByCategory[category] = append(commandsByCategory[category], builtCommand)
	}

	builtCommands := []*cli.Command{}

	for category, parentCommand := range ParentCommands {
		commands := commandsByCategory[category]

		if len(commands) > 0 {
			parentCommand.Subcommands = commands
		}

		if category == CategoryRoot {
			builtCommands = append(builtCommands, commands...)
			continue
		}

		builtCommands = append(builtCommands, parentCommand)
	}

	return builtCommands
}
