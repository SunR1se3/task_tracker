package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/constants"
	"task_tracker/internal/response"
)

func (h *Handler) MainPage(c *fiber.Ctx) error {
	return response.RenderPage(c, fiber.Map{}, "pages/main", constants.DefaultLayout)
}

func (h *Handler) AuthPage(c *fiber.Ctx) error {
	return response.RenderPage(c, nil, "pages/auth", constants.UnauthLayout)
}

func (h *Handler) UserSettingsPage(c *fiber.Ctx) error {
	return response.RenderPage(c, fiber.Map{}, "pages/user/user_profile", constants.DefaultLayout)
}
