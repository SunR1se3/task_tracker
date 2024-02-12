package response

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"task_tracker/internal/constants"
	"task_tracker/internal/errors"
	"task_tracker/internal/helper"
	"task_tracker/internal/middleware"
	"unicode/utf8"
)

type JSONResponse struct {
	Status bool     `json:"status"`
	Errors []string `json:"errors"`
	Data   any      `json:"data"`
}

func GetResponse(c *fiber.Ctx, errorHandlerInterface errors.ErrorHandlerInterface, data any) error {
	errs := []string{}
	if errorHandlerInterface != nil {
		errs = errorHandlerInterface.Get()
	}
	return c.JSON(JSONResponse{
		Status: !(len(errs) > 0),
		Errors: errs,
		Data:   data,
	})
}

func RenderPage(c *fiber.Ctx, data fiber.Map, page, layout string) error {
	if data != nil {
		claims := middleware.GetTokenClaims(c)
		firstname, lastname, _ := helper.GetUserFIO(claims)
		firstnameFl, _ := utf8.DecodeRuneInString(firstname)
		lastnameFl, _ := utf8.DecodeRuneInString(lastname)
		role := middleware.GetUserRole(c)
		data["userLogin"] = claims["login"].(string)
		data["firstname"] = firstname
		data["lastname"] = lastname
		data["headers"] = constants.Headers[role]
		data["avatarText"] = strings.ToUpper(string(firstnameFl)) + strings.ToUpper(string(lastnameFl))
	}

	return c.Render(page, data, layout)
}
