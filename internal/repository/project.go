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

func (r *ProjectRepository) CreateProject(data *domain.Project) error {
	sql := fmt.Sprintf("INSERT INTO %s(id, title, description, consumer) VALUES($1, $2, $3, $4)", constants.ProjectTable)
	_, err := r.db.Exec(sql, data.Id, data.Title, data.Description, data.Consumer)
	return err
}

func (r *ProjectRepository) GetProjectById(id uuid.UUID) (*domain.Project, error) {
	data := new(domain.Project)
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", constants.ProjectTable)
	err := r.db.Get(data, sql, id)
	return data, err
}
