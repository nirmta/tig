package mappers

import (
	OrganizationEntity "tig/app/domain/entities/organization"
	RepositoryEntity "tig/app/domain/entities/repository"
	UserEntity "tig/app/domain/entities/user"
)

type OrganizationRecord struct {
	Users []UserRecord `json:"users"`
}

type RepositoryRecord struct {
	Alias  string `json:"alias"`
	Origin string `json:"origin"`
}

type UserRecord struct {
	Alias        string             `json:"alias"`
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	Host         string             `json:"host"`
	IdentityFile string             `json:"identityFile"`
	Repositories []RepositoryRecord `json:"repositories"`
}

func OrganizationToRecordMapper(organization OrganizationEntity.Organization) OrganizationRecord {
	users := []UserRecord{}

	for _, user := range organization.GetUsers() {
		repositories := []RepositoryRecord{}
		for _, repository := range user.Repositories() {
			repositories = append(repositories, RepositoryRecord{
				Alias:  repository.Alias(),
				Origin: repository.Origin(),
			})
		}

		userRecord := UserRecord{
			Alias:        user.Alias(),
			Name:         user.Name(),
			Email:        user.Email(),
			Host:         user.Host(),
			IdentityFile: user.IdentityFile(),
			Repositories: repositories,
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

		for _, repositoryRecord := range userRecord.Repositories {
			repository := RepositoryEntity.Constructor(RepositoryEntity.RepositoryDraft{
				Alias:  repositoryRecord.Alias,
				Origin: repositoryRecord.Origin,
			})

			user.AddRepository(repository)
		}

		users = append(users, user)
	}

	org := OrganizationEntity.Constructor()
	org.SetUsers(users)
	return org
}
