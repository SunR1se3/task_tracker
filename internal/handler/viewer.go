package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/constants"
	"task_tracker/internal/middleware"
)

func (h *Handler) MainPage(c *fiber.Ctx) error {
	claims := middleware.GetTokenClaims(c)
	role := middleware.GetUserRole(c)
	return c.Render("pages/main", fiber.Map{
		"userLogin": claims["login"].(string),
		"firstname": claims["firstname"].(string),
		"lastname":  claims["lastname"].(string),
		"headers":   constants.Headers[role],
	}, "layouts/index")
}

func (h *Handler) AuthPage(c *fiber.Ctx) error {
	return c.Render("pages/auth", fiber.Map{}, "layouts/unauth_layout")
}
