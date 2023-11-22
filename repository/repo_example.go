package repository

import (
	"api-with-sql-nativ/model"
	"database/sql"
)

type RepoExample interface {
	CreateExampleData(model.ExampeModel) (model.ExampeModel, error)
	GetDataExample() ([]model.ExampeModel, error)
}

type repoExample struct {
	db *sql.DB
}

func (r *repoExample) CreateExampleData(repo model.ExampeModel) (model.ExampeModel, error) {
	err := r.db.QueryRow("INSERT INTO example VALUES($1) RETURNING example_field", repo.ExmapeData).Scan(&repo.ExmapeData)

	if err != nil {
		return model.ExampeModel{}, err
	}

	return repo, nil
}

func (r *repoExample) GetDataExample() (payload []model.ExampeModel, err error) {
	rows, err := r.db.Query("SELECT * FROM example")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		example := model.ExampeModel{}

		err = rows.Scan(&example.ExmapeData)

		if err != nil {
			return
		}

		payload = append(payload, example)
	}

	return
}

func NewRepoExample(db *sql.DB) RepoExample {
	return &repoExample{db}
}
