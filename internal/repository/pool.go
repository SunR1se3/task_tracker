package repository

import "github.com/jmoiron/sqlx"

type PoolRepository struct {
	db *sqlx.DB
}

func NewPoolRepository(db *sqlx.DB) *PoolRepository {
	return &PoolRepository{db: db}
}

//func (r *PoolRepository) GetProjectPool
