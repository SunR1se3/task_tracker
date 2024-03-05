package domain

import (
	"github.com/google/uuid"
	"task_tracker/internal/errors"
	"task_tracker/internal/helper"
	"time"
)

type Project struct {
	Id          uuid.UUID  `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Consumer    *string    `json:"consumer" db:"consumer"`
	Owner       uuid.UUID  `json:"owner"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
}

type ProjectCreateForm struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Consumer    *string `json:"consumer"`
}

type ProjectEditForm struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Consumer    *string `json:"consumer"`
}

type AddUserToTeamForm struct {
	UserId        *uuid.UUID `json:"userId"`
	ProjectId     *uuid.UUID `json:"projectId"`
	ProjectRoleId *uuid.UUID `json:"projectRoleId"`
}

type ProjectRole struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
}

type Teammate struct {
	Id             uuid.UUID    `json:"id" db:"id"`
	Fio            string       `json:"fio" db:"fio"`
	Specialization string       `json:"specialization" db:"specialization"`
	ProjectRole    *ProjectRole `json:"projectRole" db:"projectRole"`
}

type ProjectPool struct {
	ProjectId uuid.UUID `json:"projectId" db:"project_id"`
	PoolId    uuid.UUID `json:"poolId" db:"pool_id"`
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
	m.CreatedAt = time.Now()
	return nil
}

func (f *ProjectCreateForm) Validate() error {
	if len(f.Title) < 2 {
		return errors.MinFieldLengthError(helper.GetJsonTag("Title", *f), 2)
	}
	if len(f.Description) < 2 {
		return errors.MinFieldLengthError(helper.GetJsonTag("Title", *f), 2)
	}
	return nil
}

func (f *ProjectEditForm) Validate() error {
	if len(f.Title) < 2 {
		return errors.MinFieldLengthError(helper.GetJsonTag("Title", *f), 2)
	}
	if len(f.Description) < 2 {
		return errors.MinFieldLengthError(helper.GetJsonTag("Title", *f), 2)
	}
	return nil
}

func (f *ProjectEditForm) Prepare(m *Project) error {
	err := f.Validate()
	if err != nil {
		return err
	}
	m.Title = f.Title
	m.Description = f.Description
	m.Consumer = f.Consumer
	updTime := time.Now()
	m.UpdatedAt = &updTime
	return nil
}
