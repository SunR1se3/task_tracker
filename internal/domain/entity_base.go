package domain

import (
	"github.com/google/uuid"
	"time"
)

type EntityBase struct {
	Id        uuid.UUID  `json:"id" db:"id"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}
