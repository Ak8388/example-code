package usecasemanager

import (
	repomanager "api-with-sql-nativ/manager/repo_manager"
	"api-with-sql-nativ/usecase"
)

type UsecaseManager interface {
	UsecaseExampleManager() usecase.UsecaseExample
}

type usecaseManager struct {
	repoManager repomanager.RepoManager
}

func (u *usecaseManager) UsecaseExampleManager() usecase.UsecaseExample {
	return usecase.NewUsecaseExample(u.repoManager.RepoExampleManager())
}

func NewUsecaseManager(repo repomanager.RepoManager) UsecaseManager {
	return &usecaseManager{repo}
}
