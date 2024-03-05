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
	UserPicker() ([]domain.UserPicker, error)
}

type Project interface {
	CreateProject(data *domain.Project, userId uuid.UUID) error
	EditProject(data *domain.Project) error
	GetProjectById(id uuid.UUID) (*domain.Project, error)
	GetProjectsUserId(userId uuid.UUID) ([]domain.Project, error)
	AddUserToTeam(userId, projectId uuid.UUID) error
	SetUserProjectRole(userId, projectId, projectRoleId *uuid.UUID) error
	GetProjectTeam(projectId uuid.UUID) ([]domain.Teammate, error)
	GetProjectRoleForUser(projectId, userId uuid.UUID) (domain.ProjectRole, error)
	GetProjectRoles() []domain.ProjectRole
	KickUserFromTeam(userId, projectId uuid.UUID) error
}

type Sprint interface {
	CreateSprint(data *domain.Sprint, projectId uuid.UUID) error
	GetSprintById(id uuid.UUID) (*domain.Sprint, error)
	GetProjectSprints(projectId uuid.UUID) ([]domain.Sprint, error)
}

type Repository struct {
	CRUD
	Position
	Department
	Specialization
	User
	Project
	Sprint
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CRUD:           NewCRUDRepository(db),
		Position:       NewPositionRepository(db),
		Department:     NewDepartmentRepository(db),
		Specialization: NewSpecializationRepository(db),
		User:           NewUserRepository(db),
		Project:        NewProjectRepository(db),
		Sprint:         NewSprintRepository(db),
	}
}
