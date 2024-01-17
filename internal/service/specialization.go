package service

import (
	"github.com/google/uuid"
	"task_tracker/internal/domain"
	"task_tracker/internal/repository"
)

type SpecializationsService struct {
	repo repository.Specialization
}

func NewSpecializationsService(r repository.Specialization) *SpecializationsService {
	return &SpecializationsService{repo: r}
}

func (s *SpecializationsService) GetUserSpecializations(userId uuid.UUID) ([]domain.Specialization, error) {
	return s.repo.GetUserSpecializations(userId)
}

func (s *SpecializationsService) GetAll() ([]domain.Specialization, error) {
	return s.repo.GetAll()
}
