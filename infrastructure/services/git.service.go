package services

import (
	"errors"
	"os/exec"
	"strings"
	"tig/application/services"
	UserEntity "tig/domain/entities/user"
	"tig/infrastructure/consts"

	"github.com/sneuder/filesystem"
)

type GitService struct {
	fsService   *FsService
	jsonService *JsonService
}

func NewGitService(fsService *FsService, jsonService *JsonService) *GitService {
	return &GitService{
		fsService:   fsService,
		jsonService: jsonService,
	}
}

func (gs *GitService) UpdateGlobalGitConfig(config services.GlobalGitConfigInput) bool {
	exec.Command("git", "config", "--global", "user.name", config.Name).Run()
	exec.Command("git", "config", "--global", "user.email", config.Email).Run()
	exec.Command("git", "config", "--global", "user.alias", config.Alias).Run()
	return true
}

func (gs *GitService) CreateSSHKey(alias, email string) bool {
	fullPathFolderAlias := gs.getFolderPathAlias(alias)

	gs.fsService.RemoveDirectory(fullPathFolderAlias)
	gs.fsService.CreateDirectory(fullPathFolderAlias)

	fullPathNameSHH := gs.GetPrivateKeyPath(alias)
	cmd := exec.Command("ssh-keygen", "-t", "rsa", "-b", "4096", "-C", email, "-f", fullPathNameSHH, "-N", "")

	if err := cmd.Run(); err != nil {
		return false
	}

	exec.Command("eval", "$(ssh-agent -s)").Run()
	exec.Command("ssh-add", fullPathNameSHH).Run()

	return true
}

func (gs *GitService) RemoveSSHKey(alias string) bool {
	fullPathFolderAlias := gs.getFolderPathAlias(alias)
	gs.fsService.RemoveDirectory(fullPathFolderAlias)
	return true
}

func (gs *GitService) GetPrivateKeyPath(alias string) string {
	fullPathFolderAlias := gs.getFolderPathAlias(alias)
	return gs.fsService.JoinPaths([]string{fullPathFolderAlias, consts.ID_RSA})
}

func (gs *GitService) GenerateSSHConfig(users []*UserEntity.User) bool {
	directives := []filesystem.FileDirective{}

	for _, user := range users {
		directive := []filesystem.FileDirective{
			{Content: "Host " + user.Alias(), Indent: 0, NewLine: true},
			{Content: "HostName " + user.RepositoryHost(), Indent: 2, NewLine: false},
			{Content: "User git", Indent: 2, NewLine: false},
			{Content: "IdentityFile " + user.IdentityFile(), Indent: 2, NewLine: false},
		}

		directives = append(directives, directive...)
	}

	fileInfo := filesystem.FileInfo{
		Name:       consts.CONFIG,
		Path:       gs.fsService.GetFolderPath([]string{consts.SSH}),
		Directives: directives,
	}

	if _, err := filesystem.BuildFile(fileInfo, true); err != nil {
		return false
	}

	return true
}

func (gs *GitService) GetGlobalSettings() (services.GlobalGitConfigOutput, error) {
	var config services.GlobalGitConfigOutput
	notFoundError := errors.New("git settings not found")

	cmd := exec.Command("git", "config", "--global", "user.alias")
	out, err := cmd.Output()

	if err != nil {
		return config, notFoundError
	}
	config.Alias = strings.TrimSpace(string(out))

	cmd = exec.Command("git", "config", "--global", "user.name")
	out, err = cmd.Output()
	if err != nil {
		return config, notFoundError
	}
	config.Name = strings.TrimSpace(string(out))

	cmd = exec.Command("git", "config", "--global", "user.email")
	out, err = cmd.Output()
	if err != nil {
		return config, notFoundError
	}
	config.Email = strings.TrimSpace(string(out))

	return config, nil
}

func (gs *GitService) GetUserSSHConfig(alias string) (services.UserSSHConfigOutput, error) {
	sshKeyPath := gs.getFolderPathAlias(alias)
	rsaPri, _ := gs.fsService.ReadTextFile(sshKeyPath, consts.ID_RSA)
	rsaPub, _ := gs.fsService.ReadTextFile(sshKeyPath, consts.ID_RSA_PUB)

	return services.UserSSHConfigOutput{
		PublicKey:  rsaPub,
		PrivateKey: rsaPri,
	}, nil
}

//

func (gs *GitService) getFolderPathAlias(alias string) string {
	return gs.fsService.GetFolderPath([]string{consts.SSH, alias})
}
