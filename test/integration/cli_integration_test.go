package integration_test

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
	"tig/app/application/usecase"
	"tig/app/domain/entities/organization"
	"tig/app/domain/entities/repository"
	"tig/app/domain/entities/user"
	"tig/app/infrastructure/consts"
	"tig/app/infrastructure/persistence/repositories"
	"tig/app/infrastructure/services"
	"tig/app/representation/cli"
	"tig/app/representation/cli/cmd"
	"tig/app/representation/presenter"

	urfave "github.com/urfave/cli/v2"
)

type TestAppHarness struct {
	App      *urfave.App
	TempDir  string
	TomlPath string
	OrgRepo  *repositories.OrganizationRepository
}

func NewTestAppHarness(t *testing.T) *TestAppHarness {
	tempDir := t.TempDir()

	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	t.Cleanup(func() {
		os.Setenv("HOME", origHome)
	})

	tomlService := services.NewTomlService()
	fsService := services.NewFsService()

	systemFolder := fsService.GetFolderPath([]string{consts.FOLDER})
	if err := fsService.CreateDirectory(systemFolder); err != nil {
		t.Fatalf("failed to create test directory: %v", err)
	}

	tomlPath := fsService.GetFolderPath([]string{consts.FOLDER, tomlService.AddExtension(consts.FILE_NAME)})
	if err := tomlService.CreateFile(organization.New(), systemFolder, consts.FILE_NAME); err != nil {
		t.Fatalf("failed to initialize test toml: %v", err)
	}

	gitService := services.NewGitService(fsService)
	orgRepo := repositories.NewOrganizationRepository(tomlService, fsService, gitService)

	newGetUsersUseCase := usecase.NewGetUsersUseCase(orgRepo, gitService)
	addNewUserUseCase := usecase.NewAddNewUserUseCase(orgRepo, gitService)
	removeUserUseCase := usecase.NewRemoveUserUseCase(orgRepo, gitService)
	checkoutUserUseCase := usecase.NewCheckoutUserUseCase(orgRepo, gitService)
	getUserGitConfigUseCase := usecase.NewGetUserGitConfigUseCase(orgRepo, gitService)
	generateKeysUserUseCase := usecase.NewGenerateKeysUserUseCase(orgRepo, gitService)

	addRepositoryUseCase := usecase.NewAddNewRepositoryUseCase(orgRepo, gitService)
	removeRepositoryUseCase := usecase.NewRemoveRepositoryUseCase(orgRepo, gitService)
	getRepositoriesUseCase := usecase.NewGetRepositoriesUseCase(orgRepo)
	cloneRepositoryUseCase := usecase.NewDownloadRepositoriesUseCase(orgRepo, gitService, fsService)

	printService := presenter.NewPrintService()

	addUserCmd := cmd.NewAddUserCmd(addNewUserUseCase)
	removeUserCmd := cmd.NewRemoveUserCmd(removeUserUseCase)
	listUserCmd := cmd.NewListUsersCmd(newGetUsersUseCase, printService)
	checkoutUserCmd := cmd.NewCheckoutUserCmd(checkoutUserUseCase)
	listUserGitConfigCmd := cmd.NewListUserGitConfigCmd(getUserGitConfigUseCase, printService)
	generateKeysUser := cmd.NewGenerateKeysUser(generateKeysUserUseCase)

	addRepositoryCmd := cmd.NewAddRepositoryCmd(addRepositoryUseCase)
	removeRepositoryCmd := cmd.NewRemoveRepositoryCmd(removeRepositoryUseCase)
	listRepositoriesCmd := cmd.NewListRepositoriesCmd(getRepositoriesUseCase, printService)
	cloneRepositoriesCmd := cmd.NewCloneRepositoriesCmd(cloneRepositoryUseCase)

	commands := cli.Build(cli.CmdCollections{
		User:       []cli.Command{addUserCmd, removeUserCmd, listUserCmd, checkoutUserCmd, listUserGitConfigCmd, generateKeysUser},
		Repository: []cli.Command{addRepositoryCmd, removeRepositoryCmd, listRepositoriesCmd, cloneRepositoriesCmd},
	})

	app := &urfave.App{
		Name:     "tig",
		Commands: commands,
		Writer:   io.Discard,
	}

	return &TestAppHarness{
		App:      app,
		TempDir:  tempDir,
		TomlPath: tomlPath,
		OrgRepo:  orgRepo,
	}
}

func TestCLI_User_Add_Success(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	args := []string{"tig", "user", "add", "-a", "alice", "-n", "Alice Smith", "-e", "alice@tig.com", "-H", "g"}

	// Act
	err := harness.App.Run(args)

	// Assert
	if err != nil {
		t.Fatalf("expected command execution to succeed, got: %v", err)
	}

	u, err := harness.OrgRepo.GetUserByAlias("alice")
	if err != nil || u == nil {
		t.Fatalf("expected user 'alice' to be persisted in organization.toml")
	}

	if u.Name() != "Alice Smith" || u.Email() != "alice@tig.com" {
		t.Errorf("persisted user fields mismatch: name=%q email=%q", u.Name(), u.Email())
	}
}

func TestCLI_User_Add_InvalidEmail(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	args := []string{"tig", "user", "add", "-a", "bob", "-n", "Bob Doe", "-e", "invalid-email-address", "-H", "g"}

	// Act
	err := harness.App.Run(args)

	// Assert
	if err == nil {
		t.Fatalf("expected command to fail with invalid email error")
	}

	u, _ := harness.OrgRepo.GetUserByAlias("bob")
	if u != nil {
		t.Fatalf("expected invalid user not to be created")
	}
}

func TestCLI_User_List(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	_ = harness.App.Run([]string{"tig", "user", "add", "-a", "alice", "-n", "Alice", "-e", "alice@tig.com", "-H", "g"})
	var buf bytes.Buffer
	harness.App.Writer = &buf

	// Act
	err := harness.App.Run([]string{"tig", "user", "list"})

	// Assert
	if err != nil {
		t.Fatalf("expected 'tig user list' to run without error: %v", err)
	}
}

func TestCLI_User_Checkout(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	_ = harness.App.Run([]string{"tig", "user", "add", "-a", "work", "-n", "Work User", "-e", "work@company.com", "-H", "g"})

	// Act
	err := harness.App.Run([]string{"tig", "user", "checkout", "-a", "work"})

	// Assert
	if err != nil {
		t.Fatalf("expected 'tig user checkout' to succeed: %v", err)
	}
}

func TestCLI_User_Checkout_NotFound(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)

	// Act
	err := harness.App.Run([]string{"tig", "user", "checkout", "-a", "non-existent"})

	// Assert
	if err != user.ErrUserNotFound {
		t.Fatalf("expected ErrUserNotFound, got: %v", err)
	}
}

func TestCLI_User_Remove(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	_ = harness.App.Run([]string{"tig", "user", "add", "-a", "temp-user", "-n", "Temp", "-e", "temp@tig.com", "-H", "g"})
	if u, _ := harness.OrgRepo.GetUserByAlias("temp-user"); u == nil {
		t.Fatalf("expected user 'temp-user' to exist before removal")
	}

	// Act
	err := harness.App.Run([]string{"tig", "user", "remove", "-a", "temp-user"})

	// Assert
	if err != nil {
		t.Fatalf("expected 'tig user remove' to succeed: %v", err)
	}

	u, _ := harness.OrgRepo.GetUserByAlias("temp-user")
	if u != nil {
		t.Fatalf("expected user 'temp-user' to be completely removed")
	}
}

func TestCLI_User_GenerateKeys(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	_ = harness.App.Run([]string{"tig", "user", "add", "-a", "keygen-user", "-n", "Keygen", "-e", "keygen@tig.com", "-H", "g"})
	_ = harness.App.Run([]string{"tig", "user", "checkout", "-a", "keygen-user"})

	// Act
	err := harness.App.Run([]string{"tig", "user", "generate"})

	// Assert
	if err != nil {
		t.Fatalf("expected 'tig user keys' to run without error: %v", err)
	}
}

func TestCLI_User_Config(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	_ = harness.App.Run([]string{"tig", "user", "add", "-a", "conf-user", "-n", "Config User", "-e", "conf@tig.com", "-H", "g"})
	_ = harness.App.Run([]string{"tig", "user", "checkout", "-a", "conf-user"})

	// Act
	err := harness.App.Run([]string{"tig", "user", "config"})

	// Assert
	if err != nil {
		t.Fatalf("expected 'tig user config' to execute: %v", err)
	}
}

func TestCLI_Repository_Add_Success(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	_ = harness.App.Run([]string{"tig", "user", "add", "-a", "lead-dev", "-n", "Lead Dev", "-e", "lead@tig.com", "-H", "g"})
	_ = harness.App.Run([]string{"tig", "user", "checkout", "-a", "lead-dev"})

	testCases := []struct {
		name           string
		alias          string
		origin         string
		expectedOrigin string
	}{
		{
			name:           "github standard scp",
			alias:          "gh-repo",
			origin:         "git@github.com:nirmta/tig.git",
			expectedOrigin: "lead-dev:nirmta/tig.git",
		},
		{
			name:           "bitbucket with ssh subdomain",
			alias:          "bb-repo",
			origin:         "git@ssh.bitbucket.org:mycompany/backend-api.git",
			expectedOrigin: "lead-dev:mycompany/backend-api.git",
		},
		{
			name:           "gitlab standard scp",
			alias:          "gl-repo",
			origin:         "git@gitlab.com:group/subgroup/project.git",
			expectedOrigin: "lead-dev:group/subgroup/project.git",
		},
		{
			name:           "ssh protocol uri syntax",
			alias:          "uri-repo",
			origin:         "ssh://git@github.com/nirmta/awesome.git",
			expectedOrigin: "lead-dev:nirmta/awesome.git",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			err := harness.App.Run([]string{"tig", "repository", "add", "-a", tc.alias, "-o", tc.origin})

			// Assert
			if err != nil {
				t.Fatalf("expected 'tig repository add' to succeed: %v", err)
			}

			u, err := harness.OrgRepo.GetUserByAlias("lead-dev")
			if err != nil || u == nil {
				t.Fatalf("failed to retrieve user")
			}

			repo := u.FindRepositoryByAlias(tc.alias)
			if repo == nil {
				t.Fatalf("expected repository '%s' to be attached to user 'lead-dev'", tc.alias)
			}

			if repo.Origin() != tc.expectedOrigin {
				t.Errorf("expected origin '%s', got '%s'", tc.expectedOrigin, repo.Origin())
			}
		})
	}
}

func TestCLI_Repository_Add_Duplicate(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	_ = harness.App.Run([]string{"tig", "user", "add", "-a", "dup-owner", "-n", "Dup", "-e", "dup@tig.com", "-H", "g"})
	_ = harness.App.Run([]string{"tig", "user", "checkout", "-a", "dup-owner"})
	_ = harness.App.Run([]string{"tig", "repository", "add", "-a", "same-repo", "-o", "git@github.com:org/same.git"})

	// Act
	err := harness.App.Run([]string{"tig", "repository", "add", "-a", "same-repo", "-o", "git@github.com:org/same.git"})

	// Assert
	if err != repository.ErrRepositoryAlreadyExists {
		t.Fatalf("expected ErrRepositoryAlreadyExists, got: %v", err)
	}
}

func TestCLI_Repository_List(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	_ = harness.App.Run([]string{"tig", "user", "add", "-a", "list-owner", "-n", "List Owner", "-e", "list@tig.com", "-H", "g"})
	_ = harness.App.Run([]string{"tig", "user", "checkout", "-a", "list-owner"})
	_ = harness.App.Run([]string{"tig", "repository", "add", "-a", "repo-1", "-o", "git@github.com:org/repo-1.git"})

	// Act
	err := harness.App.Run([]string{"tig", "repository", "list"})

	// Assert
	if err != nil {
		t.Fatalf("expected 'tig repository list' to succeed: %v", err)
	}
}

func TestCLI_Repository_Remove(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	_ = harness.App.Run([]string{"tig", "user", "add", "-a", "rem-owner", "-n", "Rem Owner", "-e", "rem@tig.com", "-H", "g"})
	_ = harness.App.Run([]string{"tig", "user", "checkout", "-a", "rem-owner"})
	_ = harness.App.Run([]string{"tig", "repository", "add", "-a", "to-remove", "-o", "git@github.com:org/to-remove.git"})

	// Act
	err := harness.App.Run([]string{"tig", "repository", "remove", "-a", "to-remove"})

	// Assert
	if err != nil {
		t.Fatalf("expected 'tig repository remove' to succeed: %v", err)
	}

	u, _ := harness.OrgRepo.GetUserByAlias("rem-owner")
	if u.FindRepositoryByAlias("to-remove") != nil {
		t.Fatalf("expected repository 'to-remove' to be deleted")
	}
}

func TestCLI_Repository_Clone(t *testing.T) {
	// Arrange
	harness := NewTestAppHarness(t)
	_ = harness.App.Run([]string{"tig", "user", "add", "-a", "clone-owner", "-n", "Clone Owner", "-e", "clone@tig.com", "-H", "g"})
	_ = harness.App.Run([]string{"tig", "user", "checkout", "-a", "clone-owner"})
	_ = harness.App.Run([]string{"tig", "repository", "add", "-a", "project", "-o", "git@github.com:nirmta/tig.git"})

	// Act
	err := harness.App.Run([]string{"tig", "repository", "clone"})

	// Assert
	_ = err
	userDir := filepath.Join(harness.TempDir, "clone-owner")
	_ = userDir
}
