package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"task_tracker/internal/constants"
	"task_tracker/internal/errors"
	"task_tracker/internal/middleware"
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

func (h *Handler) ProjectsPage(c *fiber.Ctx) error {
	userId := middleware.GetUserId(c)
	projects, err := h.services.Project.GetProjectsUserId(*userId)
	if err != nil {
		return response.RenderPage(c, fiber.Map{}, "pages/project/projects_page", constants.DefaultLayout)
	}
	return response.RenderPage(c, fiber.Map{
		"projects": projects,
	}, "pages/project/projects_page", constants.DefaultLayout)
}

func (h *Handler) ProjectSettingsPage(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	id, err := uuid.Parse(c.Params(constants.ParamId))
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	project, err := h.services.Project.GetProjectById(id)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	projectTeam, err := h.services.Project.GetProjectTeam(project.Id)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	projectRoles := h.services.Project.GetProjectRoles()
	return response.RenderPage(c, fiber.Map{
		"projectTeam":  projectTeam,
		"project":      project,
		"projectRoles": projectRoles,
	}, "pages/project/project_settings", constants.DefaultLayout)
}
