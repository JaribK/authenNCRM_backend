package handlers

import (
	"authenncrm/internal/entities"
	"authenncrm/internal/services"
	"authenncrm/pkg/responses"

	"github.com/gofiber/fiber/v3"
)

type (
	AuthenticationHandler interface {
		Login(c fiber.Ctx) error
	}

	authenticationHandler struct {
		authenticationService services.AuthenticationService
	}
)

func NewAuthenticationHandler(authenticationService services.AuthenticationService) AuthenticationHandler {
	return &authenticationHandler{authenticationService: authenticationService}
}

func (h *authenticationHandler) Login(c fiber.Ctx) error {
	var request entities.LoginRequest

	if err := c.Bind().Body(&request); err != nil {
		return responses.BuildErrorResponse(c, fiber.StatusUnauthorized, err.Error(), "-")
	}

	response, err := h.authenticationService.Login(request)
	if err != nil {
		return responses.BuildErrorResponse(c, fiber.StatusUnauthorized, err.Error(), "-")
	}

	return responses.BuildResponse(c, fiber.StatusOK, "Login Success", response, "-")
}
