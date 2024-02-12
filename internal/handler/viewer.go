package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/constants"
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

func (h *Handler) ProjectsPages(c *fiber.Ctx) error {
	userId := middleware.GetUserId(c)
	projects, err := h.services.Project.GetProjectsUserId(*userId)
	if err != nil {
		return response.RenderPage(c, fiber.Map{}, "pages/project/projects_page", constants.DefaultLayout)
	}
	return response.RenderPage(c, fiber.Map{
		"projects": projects,
	}, "pages/project/projects_page", constants.DefaultLayout)
}
