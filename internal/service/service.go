package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"task_tracker/internal/domain"
	"task_tracker/internal/repository"
)

var Services *Service

type Position interface {
	GetUserPositions(userId uuid.UUID) ([]domain.Position, error)
	GetAll() ([]domain.Position, error)
}

type Department interface {
	CreateDepartment(formData *domain.DepartmentCreateForm) (*uuid.UUID, error)
	GetUserDepartments(userId uuid.UUID) ([]domain.Department, error)
	GetAll() ([]domain.Department, error)
}

type Specialization interface {
	GetUserSpecializations(userId uuid.UUID) ([]domain.Specialization, error)
	GetAll() ([]domain.Specialization, error)
}

type CRUD interface {
	Create(data any, tableName string) error
}

type User interface {
	CreateUser(formData *domain.UserCreateForm) (*uuid.UUID, error)
	GetUserById(id uuid.UUID) (*domain.User, error)
	GetUserDTOById(id uuid.UUID) (*domain.UserDTO, error)
	GetUsersDTO() ([]domain.UserDTO, error)
	AdminUsersTable() (*string, error)
	ChangePassword(formData *domain.ChangePasswordForm, userId *uuid.UUID) []error
	GetEditUserModalForm(id uuid.UUID) (*string, error)
	EditUser(id uuid.UUID, formData *domain.UserEditForm) error
	DisableUser(userId uuid.UUID, disable bool) error
}

type Auth interface {
	Auth(formData *domain.AuthForm) (*jwt.Token, error)
}

type Project interface {
	CreateProject(formData *domain.ProjectCreateForm, userId uuid.UUID) (*uuid.UUID, error)
	GetProjectById(id uuid.UUID) (*domain.Project, error)
	GetProjectsUserId(userId uuid.UUID) ([]domain.Project, error)
}

type Service struct {
	CRUD
	Department
	Position
	Specialization
	User
	Auth
	Project
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		CRUD:           NewCRUDService(r.CRUD),
		Position:       NewPositionService(r.Position),
		Department:     NewDepartmentService(r.Department),
		Specialization: NewSpecializationsService(r.Specialization),
		User:           NewUserService(r.User),
		Auth:           NewAuthService(r.User),
		Project:        NewProjectService(r.Project),
	}
}
