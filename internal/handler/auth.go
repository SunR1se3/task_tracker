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
		SameSite: "lax",
	})
	c.Set("Authorization", "Bearer "+t)
	return c.RedirectToRoute("user.Test", fiber.Map{
		"Title": "Hello, World!",
	})
	//return c.RedirectToRoute("api/user/test", fiber.Map{
	//	"Title": "Hello, World!",
	//})
	//return c.JSON(map[string]interface{}{
	//	"token": t,
	//	"err":   err,
	//})
}
