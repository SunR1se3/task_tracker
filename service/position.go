package service

import (
	"github.com/google/uuid"
	"task_tracker/domain"
	"task_tracker/repository"
)

type PositionService struct {
	repo repository.Position
}

func NewPositionService(r repository.Position) *PositionService {
	return &PositionService{repo: r}
}

func (s *PositionService) GetUserPositions(userId uuid.UUID) ([]domain.Position, error) {
	return s.repo.GetUserPositions(userId)
}
