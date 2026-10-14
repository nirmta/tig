package mappers

import (
	OrganizationEntity "tig/domain/entities/organization"
	UserEntity "tig/domain/entities/user"
)

type OrganizationRecord struct {
	Users []UserRecord `json:"users"`
}

type UserRecord struct {
	Alias        string `json:"alias"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Host         string `json:"host"`
	IdentityFile string `json:"identityFile"`
}

func OrganizationToRecordMapper(organization OrganizationEntity.Organization) OrganizationRecord {
	users := []UserRecord{}

	for _, user := range organization.GetUsers() {
		userRecord := UserRecord{
			Alias:        user.Alias(),
			Name:         user.Name(),
			Email:        user.Email(),
			Host:         user.Host(),
			IdentityFile: user.IdentityFile(),
		}

		users = append(users, userRecord)
	}

	return OrganizationRecord{
		Users: users,
	}
}

func RecordToOrganizationMapper(organizationRecord OrganizationRecord) *OrganizationEntity.Organization {
	users := []*UserEntity.User{}

	for _, userRecord := range organizationRecord.Users {
		user := UserEntity.Constructor(UserEntity.UserDraft{
			Alias:        userRecord.Alias,
			Name:         userRecord.Name,
			Email:        userRecord.Email,
			Host:         userRecord.Host,
			IdentityFile: userRecord.IdentityFile,
		})

		users = append(users, user)
	}

	org := OrganizationEntity.Constructor()
	org.SetUsers(users)
	return org
}
