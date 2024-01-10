package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/domain"
)

func (h *Handler) CreateDepartment(c *fiber.Ctx) error {
	formData := new(domain.DepartmentCreateForm)
	err := c.BodyParser(formData)
	if err != nil {
		return c.JSON(err)
	}
	id, err := h.services.Department.CreateDepartment(formData)
	return c.JSON(map[string]interface{}{
		"id":  id,
		"err": err,
	})
}
