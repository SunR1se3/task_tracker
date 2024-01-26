package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/constants"
	"task_tracker/internal/domain"
	"task_tracker/internal/errors"
	"task_tracker/internal/response"
)

func (h *Handler) CreatePosition(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	formData := new(domain.PositionCreateForm)
	err := c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	data := new(domain.Position)
	formData.Prepare(data)
	err = h.services.CRUD.Create(data, constants.PositionTable)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, nil)
}
