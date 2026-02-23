package handlers

import "github.com/gofiber/fiber/v3"

type (
	HealthCheckHandler interface {
		CheckHealth(c fiber.Ctx) error
	}

	healthCheckHandler struct{}
)

func NewHealthCheckHandler() HealthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) CheckHealth(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "server is running.",
	})
}
