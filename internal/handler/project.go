package handler

import (
	"github.com/gofiber/fiber/v2"
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
