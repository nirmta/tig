package usecases

import "tig/app/domain/constants"

type GetRepositoryHostsUseCase struct {
}

func NewGetRepositoryHostsUseCase() *GetRepositoryHostsUseCase {
	return &GetRepositoryHostsUseCase{}
}

func (grh *GetRepositoryHostsUseCase) Execute() map[string]string {
	return constants.RepositoryHosts
}
