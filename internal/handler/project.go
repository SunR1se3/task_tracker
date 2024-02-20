package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"task_tracker/internal/constants"
	"task_tracker/internal/domain"
	"task_tracker/internal/errors"
	"task_tracker/internal/middleware"
	"task_tracker/internal/response"
)

func (h *Handler) CreateProject(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	formData := new(domain.ProjectCreateForm)
	err := c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	userId := middleware.GetUserId(c)
	id, err := h.services.Project.CreateProject(formData, *userId)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, id)
}

func (h *Handler) EditProject(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	id, err := uuid.Parse(c.Params(constants.ParamId))
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	formData := new(domain.ProjectEditForm)
	err = c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	err = h.services.Project.EditProject(formData, id)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, nil)
}

func (h *Handler) GetMyProjects(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	userId := middleware.GetUserId(c)
	data, err := h.services.Project.GetProjectsUserId(*userId)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, data)
}

func (h *Handler) AddUserToTeam(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	formData := new(domain.AddUserToTeamForm)
	err := c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	err = h.services.Project.AddUserToTeam(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, nil)
}

func (h *Handler) GetProjectRoles(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	data := h.services.Project.GetProjectRoles()
	return response.GetResponse(c, errorHandler, data)
}

func (h *Handler) SetUserProjectRole(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	formData := new(domain.AddUserToTeamForm)
	err := c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	err = h.services.Project.SetUserProjectRole(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, nil)
}

func (h *Handler) GetProjectTeam(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	projectId, err := uuid.Parse(c.Params(constants.ParamId))
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	data, err := h.services.Project.GetProjectTeam(projectId)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, data)
}

func (h *Handler) KickUserFromTeam(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	formData := new(domain.AddUserToTeamForm)
	err := c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	err = h.services.Project.KickUserFromTeam(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, nil)
}
