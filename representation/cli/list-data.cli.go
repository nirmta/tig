package cli

import (
	"fmt"
	usecases "tig/application/use-cases"
	"tig/representation/mappers"
	"tig/representation/services"

	"github.com/urfave/cli/v2"
)

type ListDataCli struct {
	getRepositoryHostsUseCase *usecases.GetRepositoryHostsUseCase
	getUsersUseCase           *usecases.GetUsersUseCase
	getUserGitConfigUseCase   *usecases.GetUserGitConfigUseCase
	printService              *services.PrintService
}

func NewListDataCli(
	getRepositoryHostsUseCase *usecases.GetRepositoryHostsUseCase,
	getUsersUseCase *usecases.GetUsersUseCase,
	getUserGitConfigUseCase *usecases.GetUserGitConfigUseCase,
	printService *services.PrintService,
) *ListDataCli {
	return &ListDataCli{
		getRepositoryHostsUseCase: getRepositoryHostsUseCase,
		getUsersUseCase:           getUsersUseCase,
		printService:              printService,
		getUserGitConfigUseCase:   getUserGitConfigUseCase,
	}
}

func (ld *ListDataCli) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Flags:   ld.BuildFlags(),
		Usage:   "list users",
		Action:  ld.BuildAction,
	}
}

func (ld *ListDataCli) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    mappers.Host,
			Aliases: []string{"r"},
			Usage:   "list repository hosts",
		},
		&cli.StringFlag{
			Name:    mappers.PublicKey,
			Aliases: []string{"p"},
			Usage:   "list public key from a user",
		},
	}
}

func (ld *ListDataCli) BuildAction(cCtx *cli.Context) error {
	listRepositoryHosts := cCtx.Bool(mappers.Host)
	publicKeyUserAlias := cCtx.String(mappers.PublicKey)

	if listRepositoryHosts {
		table := [][]string{
			{"Repository Host", "Usage Key"},
		}

		repositoryHosts := ld.getRepositoryHostsUseCase.Execute()

		for key, repositoryHost := range repositoryHosts {
			table = append(table, []string{repositoryHost, key})
		}

		ld.printService.RenderTable(table)

		return nil
	}

	if publicKeyUserAlias != "" {
		gitConfig, _ := ld.getUserGitConfigUseCase.Execute(publicKeyUserAlias)
		fmt.Println(gitConfig.PublicKey)
		return nil
	}

	users := ld.getUsersUseCase.Execute()

	table := [][]string{
		{"User Name", "Email", "Alias", "Repository"},
	}

	for _, user := range users {
		name := user.Name()

		if isGlobal := user.InGlobal(); isGlobal {
			name = "*" + name
		}

		table = append(table, []string{name, user.Email(), user.Alias(), user.RepositoryHost()})
	}

	ld.printService.RenderTable(table)

	return nil
}
