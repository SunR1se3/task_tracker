package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/constants"
	"task_tracker/internal/domain"
	"task_tracker/internal/errors"
	"task_tracker/internal/response"
)

func (h *Handler) CreateSpecialization(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	formData := new(domain.SpecializationCreateForm)
	err := c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	data := new(domain.Specialization)
	formData.Prepare(data)
	err = h.services.CRUD.Create(data, constants.SpecializationTable)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, nil)
}
