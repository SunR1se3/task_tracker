package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"task_tracker/internal/constants"
	"task_tracker/internal/domain"
)

type DepartmentRepository struct {
	db *sqlx.DB
}

func NewDepartmentRepository(db *sqlx.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (r *DepartmentRepository) CreateDepartment(data *domain.Department) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("INSERT INTO %s(id, title, created_at, updated_at) VALUES($1, $2, $3, $4)", constants.DepartmentTable)
	_, err = tx.Exec(sql, data.Id, data.Title, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, chief := range data.Chiefs {
		sql = fmt.Sprintf("INSERT INTO %s(department_id, user_id) VALUES($1, $2)", constants.DepartmentChief)
		_, err = tx.Exec(sql, data.Id, chief)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for _, curator := range data.Curators {
		sql = fmt.Sprintf("INSERT INTO %s(department_id, user_id) VALUES($1, $2)", constants.DepartmentCurator)
		_, err = tx.Exec(sql, data.Id, curator)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	return err
}

func (r *DepartmentRepository) GetUserDepartments(userId uuid.UUID) ([]domain.Department, error) {
	data := []domain.Department{}
	sql := fmt.Sprintf("SELECT d.id, d.title, d.created_at FROM %s ud "+
		"LEFT JOIN departments d ON d.id = ud.department_id "+
		"WHERE ud.user_id = $1", constants.UserDepartmentTable)
	err := r.db.Select(&data, sql, userId)
	return data, err
}
