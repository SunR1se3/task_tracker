package domain

import (
	"github.com/google/uuid"
	"time"
)

type Department struct {
	Id        uuid.UUID   `json:"id" db:"id"`
	Title     string      `json:"title" db:"title"`
	Chiefs    []uuid.UUID `json:"chiefs" db:"chiefs"`
	Curators  []uuid.UUID `json:"curators" db:"curators"`
	CreatedAt time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time  `json:"updatedAt" db:"updated_at"`
}

type DepartmentShortObjDTO struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
}

type DepartmentCreateForm struct {
	Title    string      `json:"title"`
	Chiefs   []uuid.UUID `json:"chiefs"`
	Curators []uuid.UUID `json:"curators"`
}

func (f *DepartmentCreateForm) Prepare(m *Department) {
	m.Id = uuid.New()
	m.Title = f.Title
	m.Chiefs = f.Chiefs
	m.Curators = f.Curators
	m.CreatedAt = time.Now()
}
