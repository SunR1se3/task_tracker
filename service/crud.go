package service

import (
	"task_tracker/helper"
	"task_tracker/repository"
)

type CRUDService struct {
	repo repository.CRUD
}

func NewCRUDService(r repository.CRUD) *CRUDService {
	return &CRUDService{repo: r}
}

func (s *CRUDService) Create(data any, tableName string) error {
	fields, values := helper.GetEntityData(data)
	return s.repo.Create(fields, values, tableName)
}
