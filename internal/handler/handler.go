package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/middleware"
	"task_tracker/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(app *fiber.App) {
	api := app.Group("api")

	app.Get("/", h.MainPage)
	app.Get("/auth", h.AuthPage)

	position := api.Group("position")
	position.Post("/", h.CreatePosition)

	specialization := api.Group("specialization")
	specialization.Post("/", h.CreateSpecialization)

	department := api.Group("department")
	department.Post("/", h.CreateDepartment)

	user := api.Group("user")
	user.Post("/", h.CreateUser)
	user.Get("/test", middleware.Aw(), h.Test)
	user.Get("/:id", h.GetUserDTOById)

	auth := api.Group("auth")
	auth.Post("/", h.Auth)
}
