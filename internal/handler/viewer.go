package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/constants"
	"task_tracker/internal/middleware"
	"task_tracker/internal/response"
)

func (h *Handler) MainPage(c *fiber.Ctx) error {
	claims := middleware.GetTokenClaims(c)
	role := middleware.GetUserRole(c)
	return response.RenderPage(c, fiber.Map{
		"userLogin": claims["login"].(string),
		"firstname": claims["firstname"].(string),
		"lastname":  claims["lastname"].(string),
		"headers":   constants.Headers[role],
	}, "pages/main", constants.DefaultLayout)
}

func (h *Handler) AuthPage(c *fiber.Ctx) error {
	return response.RenderPage(c, nil, "pages/auth", constants.UnauthLayout)
}
