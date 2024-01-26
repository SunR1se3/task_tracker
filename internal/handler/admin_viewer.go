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
	departments, err := h.services.Department.GetAll()
	if err != nil {
		return err
	}
	specializations, err := h.services.Specialization.GetAll()
	if err != nil {
		return err
	}
	positions, err := h.services.Position.GetAll()
	if err != nil {
		return err
	}
	return c.Render("admin_pages/users/users_page", fiber.Map{
		"users":           users,
		"departments":     departments,
		"specializations": specializations,
		"positions":       positions,
		"headers":         constants.AdminHeader,
	}, "layouts/index")
}

func (h *Handler) UpdateTableUsers(c *fiber.Ctx) error {
	usersTable, err := h.services.User.AdminUsersTable()
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(map[string]interface{}{
		"data": usersTable,
	})
}
