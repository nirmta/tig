package repositories

import (
	"tig/app/domain/entities/organization"
	"tig/app/domain/entities/user"
	"tig/app/infrastructure/consts"
	"tig/app/infrastructure/persistence/mappers"
	"tig/app/infrastructure/services"
)

type OrganizationRepository struct {
	tomlService *services.TomlService
	fsService   *services.FsService
	gitService  *services.GitService
}

func NewOrganizationRepository(tomlService *services.TomlService, fsService *services.FsService, gitService *services.GitService) *OrganizationRepository {
	return &OrganizationRepository{
		tomlService: tomlService,
		fsService:   fsService,
		gitService:  gitService,
	}
}

func (or *OrganizationRepository) GetUnique() (*organization.Organization, error) {
	folderPath := or.fsService.GetFolderPath([]string{consts.FOLDER})
	data, err := or.tomlService.ReadFile(folderPath, consts.FILE_NAME)
	if err != nil {
		return organization.New(), nil
	}

	var orgRecord mappers.OrganizationRecord
	if err := or.tomlService.ParseFile(data, &orgRecord); err != nil {
		return organization.New(), nil
	}

	return mappers.RecordToOrganizationMapper(orgRecord), nil
}

func (or *OrganizationRepository) Update(org *organization.Organization) error {
	folderPath := or.fsService.GetFolderPath([]string{consts.FOLDER})
	orgRecord := mappers.OrganizationToRecordMapper(*org)
	return or.tomlService.CreateFile(orgRecord, folderPath, consts.FILE_NAME)
}

func (or *OrganizationRepository) GetUserByAlias(alias string) (*user.User, error) {
	org, err := or.GetUnique()
	if err != nil {
		return nil, err
	}

	u := org.FindUserByAlias(alias)
	if u == nil {
		return nil, user.ErrUserNotFound
	}

	return u, nil
}

func (or *OrganizationRepository) GetUserByEmail(email string) (*user.User, error) {
	org, err := or.GetUnique()
	if err != nil {
		return nil, err
	}

	for _, u := range org.GetUsers() {
		if u.Email() == email {
			return u, nil
		}
	}

	return nil, user.ErrUserNotFound
}

func (or *OrganizationRepository) GetAllUsers() ([]*user.User, error) {
	org, err := or.GetUnique()
	if err != nil {
		return nil, err
	}
	return org.GetUsers(), nil
}

func (or *OrganizationRepository) CreateUser(u *user.User) error {
	org, err := or.GetUnique()
	if err != nil {
		return err
	}

	org.AddUser(u)
	return or.Update(org)
}

func (or *OrganizationRepository) UpdateUser(u *user.User) error {
	org, err := or.GetUnique()
	if err != nil {
		return err
	}

	org.UpdateUser(u)
	return or.Update(org)
}

func (or *OrganizationRepository) GetCurrentUser() (*user.User, error) {
	globalSettings, err := or.gitService.GetGlobalSettings()
	if err != nil {
		return nil, err
	}

	u, err := or.GetUserByEmail(globalSettings.Email)
	if err != nil {
		return nil, err
	}
	return u, nil
}
