package service

import (
	"github.com/google/uuid"
	"task_tracker/internal/domain"
	"task_tracker/internal/repository"
)

type ProjectService struct {
	repo repository.Project
}

func NewProjectService(r repository.Project) *ProjectService {
	return &ProjectService{repo: r}
}

func (s *ProjectService) CreateProject(formData *domain.ProjectCreateForm) (*uuid.UUID, error) {
	data := new(domain.Project)
	err := formData.Prepare(data)
	if err != nil {
		return nil, err
	}
	err = s.repo.CreateProject(data)
	if err != nil {
		return nil, err
	}
	return &data.Id, nil
}
