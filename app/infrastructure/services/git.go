package services

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"tig/app/application/port"
	"tig/app/domain/entities/user"
	"tig/app/infrastructure/consts"

	"github.com/sneuder/filesystem"
)

type GitService struct {
	fsService *FsService
}

func NewGitService(fsService *FsService) *GitService {
	return &GitService{
		fsService: fsService,
	}
}

func (gs *GitService) UpdateGlobalGitConfig(config port.GlobalGitConfigInput) error {
	if err := exec.Command("git", "config", "--global", "user.name", config.Name).Run(); err != nil {
		return fmt.Errorf("failed to update git user.name: %w", err)
	}
	if err := exec.Command("git", "config", "--global", "user.email", config.Email).Run(); err != nil {
		return fmt.Errorf("failed to update git user.email: %w", err)
	}
	if err := exec.Command("git", "config", "--global", "user.alias", config.Alias).Run(); err != nil {
		return fmt.Errorf("failed to update git user.alias: %w", err)
	}
	return nil
}

func (gs *GitService) CreateSSHKey(alias, email string) error {
	fullPathFolderAlias := gs.getFolderPathAlias(alias)

	gs.fsService.RemoveDirectory(fullPathFolderAlias)
	if err := gs.fsService.CreateDirectory(fullPathFolderAlias); err != nil {
		return fmt.Errorf("failed to create ssh key directory: %w", err)
	}

	fullPathNameSHH := gs.GetPrivateKeyPath(alias)
	cmd := exec.Command("ssh-keygen", "-t", "rsa", "-b", "4096", "-C", email, "-f", fullPathNameSHH, "-N", "")

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ssh-keygen: %w", err)
	}

	exec.Command("eval", "$(ssh-agent -s)").Run()
	exec.Command("ssh-add", fullPathNameSHH).Run()

	return nil
}

func (gs *GitService) RemoveSSHKey(alias string) error {
	fullPathFolderAlias := gs.getFolderPathAlias(alias)
	if err := gs.fsService.RemoveDirectory(fullPathFolderAlias); err != nil {
		return fmt.Errorf("failed to remove ssh key directory: %w", err)
	}
	return nil
}

func (gs *GitService) GetPrivateKeyPath(alias string) string {
	fullPathFolderAlias := gs.getFolderPathAlias(alias)
	return gs.fsService.JoinPaths([]string{fullPathFolderAlias, consts.ID_RSA})
}

func (gs *GitService) GenerateSSHConfig(users []*user.User) error {
	directives := []filesystem.FileDirective{}

	for _, u := range users {
		directive := []filesystem.FileDirective{
			{Content: "Host " + u.Alias(), Indent: 0, NewLine: true},
			{Content: "HostName " + u.RepositoryHost(), Indent: 2, NewLine: false},
			{Content: "User git", Indent: 2, NewLine: false},
			{Content: "AddKeysToAgent yes", Indent: 2, NewLine: false},
			{Content: "IdentityFile " + u.IdentityFile(), Indent: 2, NewLine: false},
		}

		directives = append(directives, directive...)
	}

	fileInfo := filesystem.FileInfo{
		Name:       consts.CONFIG,
		Path:       gs.fsService.GetFolderPath([]string{consts.SSH}),
		Directives: directives,
	}

	if _, err := filesystem.BuildFile(fileInfo, true); err != nil {
		return fmt.Errorf("failed to build ssh config file: %w", err)
	}

	return nil
}

func (gs *GitService) GetGlobalSettings() (port.GlobalGitConfigOutput, error) {
	globalSettings := port.GlobalGitConfigOutput{}

	notFoundError := errors.New("git settings not found")

	rawUserName, err := exec.Command("git", "config", "--global", "user.name").Output()
	if err != nil {
		return globalSettings, notFoundError
	}

	rawEmail, err := exec.Command("git", "config", "--global", "user.email").Output()
	if err != nil {
		return globalSettings, notFoundError
	}

	rawAlias, _ := exec.Command("git", "config", "--global", "user.alias").Output()

	userName := strings.TrimSpace(string(rawUserName))
	email := strings.TrimSpace(string(rawEmail))
	alias := strings.TrimSpace(string(rawAlias))

	globalSettings.Alias = alias
	globalSettings.Name = userName
	globalSettings.Email = email

	return globalSettings, nil
}

func (gs *GitService) GetUserSSHConfig(alias string) (port.UserSSHConfigOutput, error) {
	userSSHConfigOutput := port.UserSSHConfigOutput{}

	privateKeyPath := gs.GetPrivateKeyPath(alias)
	publicKeyPath := gs.fsService.JoinPaths([]string{gs.getFolderPathAlias(alias), consts.ID_RSA_PUB})

	privateKey, err := gs.fsService.ReadFile(privateKeyPath)
	if err != nil {
		return userSSHConfigOutput, err
	}

	publicKey, err := gs.fsService.ReadFile(publicKeyPath)
	if err != nil {
		return userSSHConfigOutput, err
	}

	userSSHConfigOutput.PrivateKey = privateKey
	userSSHConfigOutput.PublicKey = publicKey

	return userSSHConfigOutput, nil
}

func (gs *GitService) CloneRepository(origin string, path string) error {
	cmd := exec.Command("git", "clone", origin, path)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone repository %s: %w", origin, err)
	}

	return nil
}

func (gs *GitService) getFolderPathAlias(alias string) string {
	homePath := gs.fsService.GetHomeDir()
	return gs.fsService.JoinPaths([]string{homePath, consts.SSH, alias})
}
