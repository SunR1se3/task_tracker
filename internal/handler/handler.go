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
	aw := middleware.Aw()

	api := app.Group("api")

	app.Get("/", h.MainPage)
	app.Get("/auth", h.AuthPage)
	app.Get("/main", aw, h.MainPage)

	admin := app.Group("admin")
	admin.Get("/users", aw, h.AdminUsersPage)
	admin.Get("/users/update_table", aw, h.UpdateTableUsers)

	position := api.Group("position")
	position.Post("/", h.CreatePosition)

	specialization := api.Group("specialization")
	specialization.Post("/", h.CreateSpecialization)

	department := api.Group("department")
	department.Post("/", h.CreateDepartment)

	user := api.Group("user")
	user.Post("/", h.CreateUser)
	user.Get("/:id", h.GetUserDTOById)

	auth := api.Group("auth")
	auth.Post("/login", h.Auth)
	auth.Get("/logout", h.Logout)
}
