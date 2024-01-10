package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/constants"
)

func (h *Handler) MainPage(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{
		"Title":   "He1llo, World!",
		"headers": constants.DefaultHeader,
	}, "layouts/main")
}

func (h *Handler) Auth(c *fiber.Ctx) error {
	return c.Render("pages/auth", fiber.Map{}, "layouts/main")
}
