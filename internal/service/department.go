package service

import (
	"github.com/google/uuid"
	"task_tracker/internal/domain"
	"task_tracker/internal/repository"
)

type DepartmentService struct {
	repo repository.Department
}

func NewDepartmentService(r repository.Department) *DepartmentService {
	return &DepartmentService{repo: r}
}

func (s *DepartmentService) CreateDepartment(formData *domain.DepartmentCreateForm) (*uuid.UUID, error) {
	data := new(domain.Department)
	formData.Prepare(data)
	err := s.repo.CreateDepartment(data)
	if err != nil {
		return nil, err
	}
	return &data.Id, nil
}

func (s *DepartmentService) GetUserDepartments(userId uuid.UUID) ([]domain.Department, error) {
	return s.repo.GetUserDepartments(userId)
}
