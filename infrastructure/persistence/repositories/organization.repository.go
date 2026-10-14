package repositories

import (
	"errors"
	OrganizationEntity "tig/domain/entities/organization"
	UserEntity "tig/domain/entities/user"
	"tig/infrastructure/consts"
	"tig/infrastructure/persistence/mappers"
	"tig/infrastructure/services"
)

type OrganizationRepository struct {
	jsonService *services.JsonService
	fsService   *services.FsService
}

func NewOrganizationRepository(jsonService *services.JsonService, fsService *services.FsService) *OrganizationRepository {
	return &OrganizationRepository{
		jsonService: jsonService,
		fsService:   fsService,
	}
}

func (or *OrganizationRepository) GetUnique() (*OrganizationEntity.Organization, error) {
	jsonFolderPath := or.fsService.GetFolderPath([]string{consts.FOLDER})
	data, _ := or.jsonService.ReadFile(jsonFolderPath, consts.FILE_NAME)
	var orgRecord mappers.OrganizationRecord
	or.jsonService.ParseFile(data, &orgRecord)
	return mappers.RecordToOrganizationMapper(orgRecord), nil
}

func (or *OrganizationRepository) Update(org *OrganizationEntity.Organization) error {
	jsonFolderPath := or.fsService.GetFolderPath([]string{consts.FOLDER})
	orgRecord := mappers.OrganizationToRecordMapper(*org)
	return or.jsonService.CreateFile(orgRecord, jsonFolderPath, consts.FILE_NAME)
}

func (or *OrganizationRepository) GetUserByAlias(alias string) (*UserEntity.User, error) {
	org, err := or.GetUnique()
	user := org.FindUserByAlias(alias)

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (or *OrganizationRepository) GetUserByEmail(email string) (*UserEntity.User, error) {
	org, _ := or.GetUnique()

	for _, user := range org.GetUsers() {
		if user.Email() == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (or *OrganizationRepository) GetAllUsers() ([]*UserEntity.User, error) {
	org, err := or.GetUnique()
	return org.GetUsers(), err
}

func (or *OrganizationRepository) CreateUser(user *UserEntity.User) error {
	org, err := or.GetUnique()
	if err != nil {
		return err
	}

	org.AddUser(user)
	return or.Update(org)
}

func (or *OrganizationRepository) UpdateUser(user *UserEntity.User) error {
	org, err := or.GetUnique()
	if err != nil {
		return err
	}

	org.UpdateUser(user)
	return or.Update(org)
}
