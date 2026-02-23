package handlers

import (
	"authenncrm/internal/entities"
	"authenncrm/internal/services"

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

	if err := c.Bind().CBOR(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Invalid request body.",
			"detail": err.Error(),
		})
	}

	claims := c.Locals("claims").(jwt.MapClaims)
	request.CreatedBy = claims["userId"].(string)

	response, err := h.userService.CreateUser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":  "Failed to create user.",
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully.",
		"detail":  response,
	})

}

func (h *userHandler) GetUserById(c fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userService.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *userHandler) GetAllUsers(c fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *userHandler) UpdateUser(c fiber.Ctx) error {
	id := c.Params("id")

	var request entities.UserRequest
	if err := c.Bind().CBOR(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.userService.UpdateUser(id, &request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func (h *userHandler) DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	if err := h.userService.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
