package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"task_tracker/internal/constants"
	"task_tracker/internal/errors"
	"task_tracker/internal/middleware"
	"task_tracker/internal/response"
)

func (h *Handler) AdminUsersPage(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	gr := middleware.IsGranted(c)
	if !gr {
		return fiber.NewError(fiber.StatusForbidden)
	}
	users, err := h.services.User.GetUsersDTO()
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	departments, err := h.services.Department.GetAll()
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	specializations, err := h.services.Specialization.GetAll()
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	positions, err := h.services.Position.GetAll()
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.RenderPage(c, fiber.Map{
		"users":           users,
		"departments":     departments,
		"specializations": specializations,
		"positions":       positions,
		"systemRoles":     constants.SystemRoles,
	}, "admin_pages/users/users_page", constants.DefaultLayout)
}

func (h *Handler) UpdateTableUsers(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	usersTable, err := h.services.User.AdminUsersTable()
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, usersTable)
}

func (h *Handler) GetEditUserModalForm(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	userId, err := uuid.Parse(c.Params(constants.ParamId))
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	userEditModalForm, err := h.services.User.GetEditUserModalForm(userId)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, userEditModalForm)
}
