package usecase

import (
	"api-with-sql-nativ/model"
	"api-with-sql-nativ/repository"
)

type UsecaseExample interface {
	CreateExampleData(model.ExampeModel) (model.ExampeModel, error)
	GetDataExample() ([]model.ExampeModel, error)
}

type usecaseExample struct {
	repo repository.RepoExample
}

func (u *usecaseExample) CreateExampleData(payload model.ExampeModel) (model.ExampeModel, error) {
	example, err := u.repo.CreateExampleData(payload)

	if err != nil {
		return model.ExampeModel{}, err
	}

	return example, nil
}

func (u *usecaseExample) GetDataExample() ([]model.ExampeModel, error) {
	example, err := u.repo.GetDataExample()

	if err != nil {
		return nil, err
	}

	return example, nil
}

func NewUsecaseExample(repo repository.RepoExample) UsecaseExample {
	return &usecaseExample{repo}
}
