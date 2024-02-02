package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/constants"
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
	app.Get("/user/profile", aw, h.UserSettingsPage)

	admin := app.Group("admin")
	admin.Get("/users", aw, h.AdminUsersPage)
	admin.Get("/users/update_table", aw, h.UpdateTableUsers)
	admin.Get("/users/edit/:"+constants.ParamId, aw, h.GetEditUserModalForm)

	position := api.Group("position")
	position.Post("/", h.CreatePosition)

	specialization := api.Group("specialization")
	specialization.Post("/", h.CreateSpecialization)

	department := api.Group("department")
	department.Post("/", h.CreateDepartment)

	user := api.Group("user")
	user.Post("/", h.CreateUser)
	user.Put("/change_password", aw, h.ChangePassword)
	user.Put("/:"+constants.ParamId+"/activation", h.DisableUser)
	user.Put("/:"+constants.ParamId, h.EditUser)
	user.Get("/:id", h.GetUserDTOById)

	auth := api.Group("auth")
	auth.Post("/login", h.Auth)
	auth.Get("/logout", h.Logout)
}
