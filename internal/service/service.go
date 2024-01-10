package service

import (
	"github.com/google/uuid"
	"task_tracker/internal/domain"
	"task_tracker/internal/repository"
)

var Services *Service

type Position interface {
	GetUserPositions(userId uuid.UUID) ([]domain.Position, error)
}

type Department interface {
	CreateDepartment(formData *domain.DepartmentCreateForm) (*uuid.UUID, error)
	GetUserDepartments(userId uuid.UUID) ([]domain.Department, error)
}

type Specialization interface {
	GetUserSpecializations(userId uuid.UUID) ([]domain.Specialization, error)
}

type CRUD interface {
	Create(data any, tableName string) error
}

type User interface {
	CreateUser(formData *domain.UserCreateForm) (*uuid.UUID, error)
	GetUserById(id uuid.UUID) (*domain.User, error)
	GetUserDTOById(id uuid.UUID) (*domain.UserDTO, error)
}

type Service struct {
	CRUD
	Department
	Position
	Specialization
	User
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		CRUD:           NewCRUDService(r.CRUD),
		Position:       NewPositionService(r.Position),
		Department:     NewDepartmentService(r.Department),
		Specialization: NewSpecializationsService(r.Specialization),
		User:           NewUserService(r.User),
	}
}
