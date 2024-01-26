package response

import (
	"github.com/gofiber/fiber/v2"
	"task_tracker/internal/errors"
)

type JSONResponse struct {
	Errors []string `json:"errors"`
	Data   any      `json:"data"`
}

func GetResponse(c *fiber.Ctx, errorHandlerInterface errors.ErrorHandlerInterface, data any) error {
	errs := []string{}
	if errorHandlerInterface != nil {
		errs = errorHandlerInterface.Get()
	}
	return c.JSON(JSONResponse{
		Errors: errs,
		Data:   data,
	})
}

func RenderPage(c *fiber.Ctx, data fiber.Map, page, layout string) error {
	return c.Render(page, data, layout)
}
