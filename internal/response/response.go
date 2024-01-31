package response

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/constants"
	"task_tracker/internal/errors"
	"task_tracker/internal/middleware"
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
		role := middleware.GetUserRole(c)
		data["userLogin"] = claims["login"].(string)
		data["firstname"] = claims["firstname"].(string)
		data["lastname"] = claims["lastname"].(string)
		data["headers"] = constants.Headers[role]
	}
	return c.Render(page, data, layout)
}
