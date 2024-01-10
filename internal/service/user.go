package service

import (
	"github.com/google/uuid"
	"task_tracker/internal/domain"
	"task_tracker/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(r repository.User) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(formData *domain.UserCreateForm) (*uuid.UUID, error) {
	data := new(domain.User)
	formData.Prepare(data)
	err := s.repo.CreateUser(data)
	if err != nil {
		return nil, err
	}
	return &data.Id, nil
}

func (s *UserService) GetUserById(id uuid.UUID) (*domain.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) GetUserDTOById(id uuid.UUID) (*domain.UserDTO, error) {
	positions, err := Services.Position.GetUserPositions(id)
	if err != nil {
		return nil, err
	}
	departments, err := Services.Department.GetUserDepartments(id)
	if err != nil {
		return nil, err
	}
	specializations, err := Services.Specialization.GetUserSpecializations(id)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.GetUserDTOById(id)
	if err != nil {
		return nil, err
	}
	user.Positions = positions
	user.Departments = departments
	user.Specializations = specializations
	return user, err
}
