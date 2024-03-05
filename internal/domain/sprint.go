package domain

import (
	"github.com/google/uuid"
	"task_tracker/internal/errors"
	"task_tracker/internal/helper"
	"time"
)

type Sprint struct {
	Id        uuid.UUID  `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	StartDate *time.Time `json:"startDate" db:"start_date"`
	EndDate   *time.Time `json:"endDate" db:"end_date"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

type SprintCreateForm struct {
	Title string `json:"title" db:"title"`
}

type SprintEditForm struct {
	Title   string     `json:"title" db:"title"`
	EndDate *time.Time `json:"endDate" db:"end_date"`
}

func (f *SprintCreateForm) Validate() error {
	if len(f.Title) == 0 {
		return errors.RequiredFiledError(helper.GetJsonTag("Title", *f))
	}
	return nil
}

func (f *SprintEditForm) Validate(startDate time.Time) error {
	if len(f.Title) == 0 {
		return errors.RequiredFiledError(helper.GetJsonTag("Title", *f))
	}
	if f.EndDate.Before(startDate) {
		return errors.EndDateSprintError(helper.GetJsonTag("EndDate", *f))
	}
	return nil
}

func (f *SprintCreateForm) Prepare(m *Sprint) error {
	err := f.Validate()
	if err != nil {
		return err
	}
	m.Id = uuid.New()
	m.Title = f.Title
	m.CreatedAt = time.Now()
	return nil
}

func (f *SprintEditForm) Prepare(m *Sprint) error {
	err := f.Validate(*m.EndDate)
	if err != nil {
		return err
	}
	m.Title = f.Title
	t := time.Now()
	m.UpdatedAt = &t
	return nil
}
