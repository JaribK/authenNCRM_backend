package handlers

import (
	"authenncrm/internal/entities"
	"authenncrm/internal/services"
	"authenncrm/pkg/responses"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type (
	UserHandler interface {
		CreateUser(c fiber.Ctx) error
		GetUserById(c fiber.Ctx) error
		GetAllUsers(c fiber.Ctx) error
		UpdateUser(c fiber.Ctx) error
		DeleteUser(c fiber.Ctx) error
	}

	userHandler struct {
		userService services.UserService
	}
)

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) CreateUser(c fiber.Ctx) error {
	var request entities.UserRequest

	if err := c.Bind().Body(&request); err != nil {
		return responses.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error(), "-")
	}

	claims := c.Locals("claims").(jwt.MapClaims)
	request.CreatedBy = claims["userId"].(string)

	response, err := h.userService.CreateUser(&request)
	if err != nil {
		return responses.BuildErrorResponse(c, fiber.StatusUnprocessableEntity, err.Error(), "-")
	}

	return responses.BuildResponse(c, fiber.StatusCreated, "User created successfully", response, "-")

}

func (h *userHandler) GetUserById(c fiber.Ctx) error {
	id := c.Params("userId")

	user, err := h.userService.GetUserById(id)
	if err != nil {
		return responses.BuildErrorResponse(c, fiber.StatusNotFound, err.Error(), "-")
	}

	return responses.BuildResponse(c, fiber.StatusOK, "User retrieved successfully", user, "-")
}

func (h *userHandler) GetAllUsers(c fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return responses.BuildErrorResponse(c, fiber.StatusInternalServerError, err.Error(), "-")
	}

	return responses.BuildResponse(c, fiber.StatusOK, "Users retrieved successfully", users, "-")
}

func (h *userHandler) UpdateUser(c fiber.Ctx) error {
	id := c.Params("userId")

	var request entities.UserRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.BuildErrorResponse(c, fiber.StatusBadRequest, err.Error(), "-")

	}

	if err := h.userService.UpdateUser(id, &request); err != nil {
		return responses.BuildErrorResponse(c, fiber.StatusInternalServerError, err.Error(), "-")
	}

	return responses.BuildResponse(c, fiber.StatusOK, "User updated successfully", request, "-")
}

func (h *userHandler) DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	if err := h.userService.DeleteUser(id); err != nil {
		return responses.BuildErrorResponse(c, fiber.StatusInternalServerError, err.Error(), "-")
	}

	return responses.BuildResponse(c, fiber.StatusNoContent, "User deleted successfully", nil, "-")
}
