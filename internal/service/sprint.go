package service

import (
	"github.com/google/uuid"
	"task_tracker/internal/domain"
	"task_tracker/internal/helper"
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

func (s *SprintService) GetProjectSprints(params domain.SprintParams) (*domain.SprintResponse, error) {
	data, err := s.repo.GetProjectSprints(params)
	if err != nil {
		return nil, err
	}
	total := 0
	if len(data) > 0 {
		total = *data[0].Total
	}
	return &domain.SprintResponse{
		Sprints: data,
		Total:   total,
	}, nil
}

func (s *SprintService) GetProjectSprintCards(params domain.SprintParams) (*string, error) {
	data, err := s.GetProjectSprints(params)
	if err != nil {
		return nil, err
	}
	return helper.HtmlRenderProcess("./views/pages/project/sprints.html", "sprint_cards", map[string]interface{}{
		"data": data,
	})
}
