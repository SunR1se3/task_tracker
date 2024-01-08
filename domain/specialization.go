package domain

import (
	"github.com/google/uuid"
	"time"
)

type Specialization struct {
	Id        uuid.UUID  `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

type SpecializationCreateForm struct {
	Title string `json:"title" db:"title"`
}

func (f *SpecializationCreateForm) Prepare(m *Specialization) {
	m.Id = uuid.New()
	m.Title = f.Title
	m.CreatedAt = time.Now()
}
