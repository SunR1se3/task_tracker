package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strconv"
	"task_tracker/internal/constants"
	"task_tracker/internal/domain"
	"task_tracker/internal/errors"
	"task_tracker/internal/middleware"
	"task_tracker/internal/response"
)

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	formData := new(domain.UserCreateForm)
	err := c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	id, err := h.services.User.CreateUser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, id)
}

func (h *Handler) GetUserDTOById(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	data, err := h.services.User.GetUserDTOById(id)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, data)
}

func (h *Handler) ChangePassword(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	formData := new(domain.ChangePasswordForm)
	err := c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	userId := middleware.GetUserId(c)
	errs := h.services.User.ChangePassword(formData, userId)
	if len(errs) > 0 {
		errorHandler.Add(errs...)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, nil)
}

func (h *Handler) EditUser(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	userId, err := uuid.Parse(c.Params(constants.ParamId))
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	formData := new(domain.UserEditForm)
	err = c.BodyParser(formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	err = h.services.User.EditUser(userId, formData)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, nil)
}

func (h *Handler) DisableUser(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	userId, err := uuid.Parse(c.Params(constants.ParamId))
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	disabled, err := strconv.ParseBool(c.Query(constants.QueryDisabledParam))
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	err = h.services.User.DisableUser(userId, disabled)
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, nil)
}

func (h *Handler) UserPicker(c *fiber.Ctx) error {
	errorHandler := new(errors.ErrorHandler)
	data, err := h.services.User.UserPicker()
	if err != nil {
		errorHandler.Add(err)
		return response.GetResponse(c, errorHandler, nil)
	}
	return response.GetResponse(c, errorHandler, data)
}
