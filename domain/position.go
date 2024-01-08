package domain

import (
	"github.com/google/uuid"
	"time"
)

type Position struct {
	Id        uuid.UUID  `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	Code      string     `json:"code" db:"code"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

type PositionCreateForm struct {
	Title string `json:"title"`
	Code  string `json:"code"`
}

func (f *PositionCreateForm) Prepare(m *Position) {
	m.Id = uuid.New()
	m.Title = f.Title
	m.Code = f.Code
	m.CreatedAt = time.Now()
}
