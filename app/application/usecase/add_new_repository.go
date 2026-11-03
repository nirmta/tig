package usecase

import (
	"fmt"
	"strings"
	appRepo "tig/app/application/repository"
	"tig/app/application/port"
	"tig/app/domain/entities/organization"
	domainRepo "tig/app/domain/entities/repository"
	"tig/app/domain/entities/user"
)

type AddNewRepositoryUseCaseInput struct {
	Alias  string
	Origin string
}

type AddNewRepositoryUseCase struct {
	organizationRepository appRepo.OrganizationRepository
	gitService             port.GitService
}

func NewAddNewRepositoryUseCase(organizationRepository appRepo.OrganizationRepository, gitService port.GitService) *AddNewRepositoryUseCase {
	return &AddNewRepositoryUseCase{
		organizationRepository: organizationRepository,
		gitService:             gitService,
	}
}

func (anr *AddNewRepositoryUseCase) Execute(input AddNewRepositoryUseCaseInput) error {
	currentUser, err := anr.organizationRepository.GetCurrentUser()
	if err != nil || currentUser == nil {
		return user.ErrUserNotFound
	}

	if foundRepository := currentUser.FindRepositoryByAlias(input.Alias); foundRepository != nil {
		return domainRepo.ErrRepositoryAlreadyExists
	}

	transformedOrigin := anr.formatOriginWithAlias(input.Origin, currentUser.Alias())

	newRepository, err := domainRepo.New(domainRepo.Draft{
		Alias:  input.Alias,
		Origin: transformedOrigin,
	})
	if err != nil {
		return err
	}

	currentUser.AddRepository(newRepository)

	org, err := anr.organizationRepository.GetUnique()
	if err != nil || org == nil {
		return organization.ErrOrganizationNotFound
	}

	org.UpdateUser(currentUser)
	if err := anr.organizationRepository.Update(org); err != nil {
		return fmt.Errorf("failed to update organization: %w", err)
	}

	return nil
}

func (anr *AddNewRepositoryUseCase) formatOriginWithAlias(origin string, alias string) string {
	cleaned := strings.TrimSpace(origin)
	if cleaned == "" {
		return origin
	}

	if strings.HasPrefix(cleaned, "ssh://git@") {
		trimmed := strings.TrimPrefix(cleaned, "ssh://git@")
		if slashIdx := strings.Index(trimmed, "/"); slashIdx != -1 {
			return fmt.Sprintf("%s:%s", alias, trimmed[slashIdx+1:])
		}
	}

	if strings.HasPrefix(cleaned, "git@") {
		colonIdx := strings.Index(cleaned, ":")
		if colonIdx != -1 {
			return fmt.Sprintf("%s:%s", alias, cleaned[colonIdx+1:])
		}
	}

	return origin
}
