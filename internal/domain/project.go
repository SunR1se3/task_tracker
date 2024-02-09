package domain

import (
	"github.com/google/uuid"
	"task_tracker/internal/errors"
	"task_tracker/internal/helper"
)

type Project struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Consumer    *string   `json:"consumer" db:"consumer"`
}

type ProjectCreateForm struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Consumer    *string `json:"consumer"`
}

func (f *ProjectCreateForm) Prepare(m *Project) error {
	err := f.Validate()
	if err != nil {
		return err
	}
	m.Id = uuid.New()
	m.Title = f.Title
	m.Description = f.Description
	m.Consumer = f.Consumer
	return nil
}

func (f *ProjectCreateForm) Validate() error {
	if len(f.Title) < 2 {
		return errors.MinFieldLengthError(helper.GetJsonTag("Title", f), 2)
	}
	if len(f.Description) < 2 {
		return errors.MinFieldLengthError(helper.GetJsonTag("Title", f), 2)
	}
	return nil
}
