package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"task_tracker/internal/constants"
	"task_tracker/internal/domain"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) CreateProject(data *domain.Project, userId uuid.UUID) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("INSERT INTO %s(id, title, description, consumer, owner, created_at) VALUES($1, $2, $3, $4, $5, $6)", constants.ProjectTable)
	_, err = tx.Exec(sql, data.Id, data.Title, data.Description, data.Consumer, userId, data.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	sql = fmt.Sprintf("INSERT INTO %s(user_id, project_id) VALUES($1, $2)", constants.UserProjectTable)
	_, err = tx.Exec(sql, userId, data.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (r *ProjectRepository) EditProject(data *domain.Project) error {
	sql := fmt.Sprintf("UPDATE %s SET title = $1, description = $2, consumer = $3, updated_at = $4 WHERE id = $5", constants.ProjectTable)
	_, err := r.db.Exec(sql, data.Title, data.Description, data.Consumer, data.UpdatedAt, data.Id)
	return err
}

func (r *ProjectRepository) GetProjectById(id uuid.UUID) (*domain.Project, error) {
	data := new(domain.Project)
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", constants.ProjectTable)
	err := r.db.Get(data, sql, id)
	return data, err
}

func (r *ProjectRepository) GetProjectsUserId(userId uuid.UUID) ([]domain.Project, error) {
	data := []domain.Project{}
	sql := fmt.Sprintf("SELECT p.* FROM %s p "+
		"LEFT JOIN %s up "+
		"ON up.project_id = p.id "+
		"WHERE up.user_id = $1", constants.ProjectTable, constants.UserProjectTable)
	err := r.db.Select(&data, sql, userId)
	return data, err
}

func (r *ProjectRepository) AddUserToTeam(userId, projectId uuid.UUID) error {
	sql := fmt.Sprintf("INSERT INTO %s(user_id, project_id) VALUES($1, $2)", constants.UserProjectTable)
	_, err := r.db.Exec(sql, userId, projectId)
	return err
}

func (r *ProjectRepository) SetUserProjectRole(userId, projectId, projectRoleId *uuid.UUID) error {
	sql := fmt.Sprintf("UPDATE %s SET project_role_id = $3 WHERE user_id = $1 and project_id = $2", constants.UserProjectTable)
	_, err := r.db.Exec(sql, userId, projectId, projectRoleId)
	return err
}

func (r *ProjectRepository) GetProjectTeam(projectId uuid.UUID) ([]domain.Teammate, error) {
	data := []domain.Teammate{}
	sql := "SELECT u.id, concat_ws(' ', u.lastname, u.firstname, u.middlename) AS fio, " +
		"s.title as specialization FROM users u " +
		"left join user_specialization us " +
		"on us.user_id = u.id " +
		"LEFT JOIN specializations s " +
		"ON s.id = us.specialization_id " +
		"LEFT JOIN user_project up " +
		"ON up.user_id = u.id and up.project_id = $1 " +
		"LEFT JOIN project_roles pr " +
		"ON pr.id = up.project_role_id " +
		"WHERE up.project_id = $1"
	err := r.db.Select(&data, sql, projectId)
	return data, err
}

func (r *ProjectRepository) GetProjectRoleForUser(projectId, userId uuid.UUID) (domain.ProjectRole, error) {
	data := domain.ProjectRole{}
	sql := fmt.Sprintf("SELECT pr.* FROM %s pr "+
		"LEFT JOIN %s up "+
		"ON up.project_role_id = pr.id "+
		"WHERE up.user_id = $1 AND up.project_id = $2", constants.UserProjectRoles, constants.UserProjectTable)
	err := r.db.Get(&data, sql, userId, projectId)
	return data, err
}

func (r *ProjectRepository) GetProjectRoles() []domain.ProjectRole {
	data := []domain.ProjectRole{}
	sql := fmt.Sprintf("SELECT * FROM %s", constants.UserProjectRoles)
	_ = r.db.Select(&data, sql)
	return data
}

func (r *ProjectRepository) KickUserFromTeam(userId, projectId uuid.UUID) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 and project_id = $2", constants.UserProjectTable)
	_, err := r.db.Exec(sql, userId, projectId)
	return err
}
