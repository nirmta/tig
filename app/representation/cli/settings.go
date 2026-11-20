package cli

import "github.com/urfave/cli/v2"

const (
	FlagAlias       = "alias"
	FlagAliasShort  = "a"
	FlagName        = "name"
	FlagNameShort   = "n"
	FlagEmail       = "email"
	FlagEmailShort  = "e"
	FlagHost        = "host"
	FlagHostShort   = "H"
	FlagOrigin      = "origin"
	FlagOriginShort = "o"
	FlagPath        = "path"
	FlagOutput      = "output"
	FlagOutputShort = "o"
	FlagInput       = "input"
	FlagInputShort  = "i"
)

type Category string

const (
	CategoryUser       Category = "user"
	CategoryRepository Category = "repository"
	CategoryRoot       Category = "root"
)

var ParentCommands = map[Category]*cli.Command{
	CategoryUser: {
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "user operations",
	},
	CategoryRepository: {
		Name:    "repository",
		Aliases: []string{"r"},
		Usage:   "repository operations",
	},
	CategoryRoot: {
		Name:    "root",
		Aliases: []string{"ro"},
		Usage:   "root operations",
	},
}
