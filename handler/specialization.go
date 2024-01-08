package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/constants"
	"task_tracker/domain"
)

func (h *Handler) CreateSpecialization(c *fiber.Ctx) error {
	formData := new(domain.SpecializationCreateForm)
	err := c.BodyParser(formData)
	if err != nil {
		return c.JSON(err)
	}
	data := new(domain.Specialization)
	formData.Prepare(data)
	err = h.services.CRUD.Create(data, constants.SpecializationTable)
	return c.JSON(err)
}
