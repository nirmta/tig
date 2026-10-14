package usecases

import "tig/domain/constants"

type GetRepositoryHostsUseCase struct {
}

func NewGetRepositoryHostsUseCase() *GetRepositoryHostsUseCase {
	return &GetRepositoryHostsUseCase{}
}

func (grh *GetRepositoryHostsUseCase) Execute() map[string]string {
	return constants.RepositoryHosts
}
