package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"task_tracker/internal/constants"
	"task_tracker/internal/domain"
)

type PositionRepository struct {
	db *sqlx.DB
}

func NewPositionRepository(db *sqlx.DB) *PositionRepository {
	return &PositionRepository{db: db}
}

func (r *PositionRepository) GetUserPositions(userId uuid.UUID) ([]domain.Position, error) {
	data := []domain.Position{}
	sql := fmt.Sprintf("SELECT p.id, p.title, p.code, p.created_at FROM %s up "+
		"LEFT JOIN positions p ON p.id = up.position_id "+
		"WHERE up.user_id = $1", constants.UserPositionTable)
	err := r.db.Select(&data, sql, userId)
	return data, err
}

func (r *PositionRepository) GetAll() ([]domain.Position, error) {
	data := []domain.Position{}
	sql := fmt.Sprintf("SELECT * FROM %s", constants.PositionTable)
	err := r.db.Select(&data, sql)
	return data, err
}
