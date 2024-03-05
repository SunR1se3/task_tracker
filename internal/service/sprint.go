package service

import (
	"github.com/google/uuid"
	"task_tracker/internal/domain"
	"task_tracker/internal/repository"
)

type SprintService struct {
	repo repository.Sprint
}

func NewSprintService(r repository.Sprint) *SprintService {
	return &SprintService{repo: r}
}

func (s *SprintService) CreateSprint(formData *domain.SprintCreateForm, projectId uuid.UUID) (*uuid.UUID, error) {
	data := new(domain.Sprint)
	err := formData.Prepare(data)
	if err != nil {
		return nil, err
	}
	err = s.repo.CreateSprint(data, projectId)
	if err != nil {
		return nil, err
	}
	return &data.Id, nil
}

func (s *SprintService) GetSprintById(id uuid.UUID) (*domain.Sprint, error) {
	return s.repo.GetSprintById(id)
}

func (s *SprintService) GetProjectSprints(projectId uuid.UUID) ([]domain.Sprint, error) {
	return s.repo.GetProjectSprints(projectId)
}
