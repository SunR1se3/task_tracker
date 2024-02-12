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
	sql := fmt.Sprintf("INSERT INTO %s(id, title, description, consumer, created_at) VALUES($1, $2, $3, $4, $5)", constants.ProjectTable)
	_, err = tx.Exec(sql, data.Id, data.Title, data.Description, data.Consumer, data.CreatedAt)
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
