package repomanager

import (
	inframanager "api-with-sql-nativ/manager/infra_manager"
	"api-with-sql-nativ/repository"
)

type RepoManager interface {
	RepoExampleManager() repository.RepoExample
}

type repoManager struct {
	infra inframanager.InfraManager
}

func (infra *repoManager) RepoExampleManager() repository.RepoExample {
	return repository.NewRepoExample(infra.infra.Connect())
}

func NewRepoManager(infra inframanager.InfraManager) RepoManager {
	return &repoManager{infra}
}
