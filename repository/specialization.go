package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"task_tracker/constants"
	"task_tracker/domain"
)

type SpecializationRepository struct {
	db *sqlx.DB
}

func NewSpecializationRepository(db *sqlx.DB) *SpecializationRepository {
	return &SpecializationRepository{db: db}
}

func (r *SpecializationRepository) GetUserSpecializations(userId uuid.UUID) ([]domain.Specialization, error) {
	data := []domain.Specialization{}
	sql := fmt.Sprintf("SELECT s.id, s.title, s.created_at FROM %s us "+
		"LEFT JOIN specializations s ON s.id = us.specialization_id "+
		"WHERE us.user_id = $1", constants.UserSpecializationTable)
	err := r.db.Select(&data, sql, userId)
	return data, err
}
