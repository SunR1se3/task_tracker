package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/constants"
	"task_tracker/domain"
)

func (h *Handler) CreatePosition(c *fiber.Ctx) error {
	formData := new(domain.PositionCreateForm)
	err := c.BodyParser(formData)
	if err != nil {
		return c.JSON(err)
	}
	data := new(domain.Position)
	formData.Prepare(data)
	err = h.services.CRUD.Create(data, constants.PositionTable)
	return c.JSON(err)
}
