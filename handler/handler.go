package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(app *fiber.App) {
	api := app.Group("api")

	position := api.Group("position")
	position.Post("/", h.CreatePosition)

	specialization := api.Group("specialization")
	specialization.Post("/", h.CreateSpecialization)

	department := api.Group("department")
	department.Post("/", h.CreateDepartment)

	user := api.Group("user")
	user.Post("/", h.CreateUser)
	user.Get("/:id", h.GetUserDTOById)
}
