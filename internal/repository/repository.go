package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"task_tracker/internal/domain"
)

type CRUD interface {
	Create(fields []string, values []any, tableName string) error
}

type Position interface {
	GetUserPositions(userId uuid.UUID) ([]domain.Position, error)
	GetAll() ([]domain.Position, error)
}

type Department interface {
	CreateDepartment(data *domain.Department) error
	GetUserDepartments(userId uuid.UUID) ([]domain.Department, error)
	GetAll() ([]domain.Department, error)
}

type Specialization interface {
	GetUserSpecializations(userId uuid.UUID) ([]domain.Specialization, error)
	GetAll() ([]domain.Specialization, error)
}

type User interface {
	CreateUser(data *domain.User) error
	GetUserById(id uuid.UUID) (*domain.User, error)
	GetUserDTOById(id uuid.UUID) (*domain.UserDTO, error)
	AlreadyExists(login string) (bool, error)
	GetUserByLogin(login string) (*domain.User, error)
	GetUsersDTO() ([]domain.UserDTO, error)
	ChangePassword(newPassword string, userId *uuid.UUID) error
	EditUser(data *domain.User) error
	DisableUser(userId uuid.UUID, disable bool) error
}

type Repository struct {
	CRUD
	Position
	Department
	Specialization
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CRUD:           NewCRUDRepository(db),
		Position:       NewPositionRepository(db),
		Department:     NewDepartmentRepository(db),
		Specialization: NewSpecializationRepository(db),
		User:           NewUserRepository(db),
	}
}
