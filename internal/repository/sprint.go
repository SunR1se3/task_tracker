package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
	"task_tracker/internal/constants"
	"task_tracker/internal/domain"
)

type SprintRepository struct {
	db *sqlx.DB
}

func NewSprintRepository(db *sqlx.DB) *SprintRepository {
	return &SprintRepository{db: db}
}

func (r *SprintRepository) CreateSprint(data *domain.Sprint, projectId uuid.UUID) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("INSERT INTO %s(id, title, created_at) VALUES($1, $2, $3)", constants.SprintsTable)
	_, err = tx.Exec(sql, data.Id, data.Title, data.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	sql = fmt.Sprintf("INSERT INTO %s(project_id, sprint_id) VALUES($1, $2)", constants.ProjectSprintsTable)
	_, err = tx.Exec(sql, projectId, data.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (r *SprintRepository) GetSprintById(id uuid.UUID) (*domain.Sprint, error) {
	data := new(domain.Sprint)
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", constants.SprintsTable)
	err := r.db.Get(data, sql, id)
	return data, err
}

func (r *SprintRepository) GetProjectSprints(params domain.SprintParams) ([]domain.Sprint, error) {
	data := []domain.Sprint{}
	sql := fmt.Sprintf("SELECT count(*) OVER() AS total, s.* FROM %s s "+
		"LEFT JOIN %s ps "+
		"ON ps.sprint_id = s.id ", constants.SprintsTable, constants.ProjectSprintsTable)
	wherePart, mapParams := r.sprintFilter(params)
	if wherePart != "" {
		sql += "WHERE " + wherePart
	}
	if *params.CreatedAtOrder == 0 {
		sql += " ORDER BY s.created_at desc"
	} else {
		sql += " ORDER BY s.created_at"
	}

	sql += " LIMIT :limit OFFSET :offset"
	query, args, err := sqlx.Named(sql, mapParams)
	query = r.db.Rebind(query)
	err = r.db.Select(&data, query, args...)
	return data, err
}

func (r *SprintRepository) sprintFilter(params domain.SprintParams) (string, map[string]interface{}) {
	whereParts := []string{}
	if params.ProjectId != nil {
		whereParts = append(whereParts, "ps.project_id = :projectId")
	}
	if params.Title != nil {
		whereParts = append(whereParts, "s.title ILIKE :title")
	}
	return strings.Join(whereParts, " and "), params.PrepareParams()
}
