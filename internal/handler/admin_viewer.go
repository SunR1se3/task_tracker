package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/constants"
	"task_tracker/internal/middleware"
)

func (h *Handler) AdminUsersPage(c *fiber.Ctx) error {
	gr := middleware.IsGranted(c)
	if !gr {
		return fiber.NewError(fiber.StatusForbidden)
	}
	users, err := h.services.User.GetUsersDTO()
	if err != nil {
		return err
	}
	return c.Render("admin_pages/users_page", fiber.Map{
		"users":   users,
		"headers": constants.AdminHeader,
	}, "layouts/index")
}
