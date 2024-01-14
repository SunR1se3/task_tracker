package handler

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/domain"
	"time"
)

func (h *Handler) Auth(c *fiber.Ctx) error {
	formData := new(domain.AuthForm)
	err := c.BodyParser(formData)
	if err != nil {
		return err
	}
	token, err := h.services.Auth.Auth(formData)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
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
	return c.JSON("main")
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
