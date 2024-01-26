package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/domain"
	"task_tracker/internal/errors"
	"task_tracker/internal/response"
	"time"
)

func (h *Handler) Auth(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	formData := new(domain.AuthForm)
	err := c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	token, err := h.services.Auth.Auth(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Domain:   "ttracker.test",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
		SameSite: "Strict",
	})
	return response.GetResponse(c, errorHandler, "main")
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Domain:   "ttracker.test",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 100),
		HTTPOnly: true,
		SameSite: "Strict",
	})
	return c.Redirect("/auth")
}
