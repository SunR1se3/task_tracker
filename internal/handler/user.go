package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"task_tracker/internal/domain"
)

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	formData := new(domain.UserCreateForm)
	err := c.BodyParser(formData)
	if err != nil {
		return c.JSON(err)
	}
	id, err := h.services.User.CreateUser(formData)
	return c.JSON(map[string]interface{}{
		"id":  id,
		"err": err,
	})
}

func (h *Handler) GetUserDTOById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}
	data, err := h.services.User.GetUserDTOById(id)
	return c.JSON(map[string]interface{}{
		"data": data,
		"err":  err,
	})
}

func (h *Handler) Test(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}
