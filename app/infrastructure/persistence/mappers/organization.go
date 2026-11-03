package mappers

import (
	"tig/app/domain/entities/organization"
	"tig/app/domain/entities/repository"
	"tig/app/domain/entities/user"
)

type OrganizationRecord struct {
	Users []UserRecord `toml:"users"`
}

type RepositoryRecord struct {
	Alias  string `toml:"alias"`
	Origin string `toml:"origin"`
}

type UserRecord struct {
	Alias        string             `toml:"alias"`
	Name         string             `toml:"name"`
	Email        string             `toml:"email"`
	Host         string             `toml:"host"`
	IdentityFile string             `toml:"identity_file"`
	Repositories []RepositoryRecord `toml:"repositories"`
}

func OrganizationToRecordMapper(org organization.Organization) OrganizationRecord {
	users := []UserRecord{}

	for _, u := range org.GetUsers() {
		repos := []RepositoryRecord{}
		for _, r := range u.Repositories() {
			repos = append(repos, RepositoryRecord{
				Alias:  r.Alias(),
				Origin: r.Origin(),
			})
		}

		userRecord := UserRecord{
			Alias:        u.Alias(),
			Name:         u.Name(),
			Email:        u.Email(),
			Host:         u.Host(),
			IdentityFile: u.IdentityFile(),
			Repositories: repos,
		}

		users = append(users, userRecord)
	}

	return OrganizationRecord{
		Users: users,
	}
}

func RecordToOrganizationMapper(organizationRecord OrganizationRecord) *organization.Organization {
	users := make([]*user.User, 0, len(organizationRecord.Users))

	for _, userRecord := range organizationRecord.Users {
		repos := make([]*repository.Repository, 0, len(userRecord.Repositories))
		for _, repositoryRecord := range userRecord.Repositories {
			repo := repository.Reconstitute(repository.Draft{
				Alias:  repositoryRecord.Alias,
				Origin: repositoryRecord.Origin,
			})
			repos = append(repos, repo)
		}

		u := user.Reconstitute(user.Draft{
			Alias:        userRecord.Alias,
			Name:         userRecord.Name,
			Email:        userRecord.Email,
			Host:         userRecord.Host,
			IdentityFile: userRecord.IdentityFile,
		}, false, repos)

		users = append(users, u)
	}

	return organization.Reconstitute(users)
}
