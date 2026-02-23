package handlers

import (
	"authenncrm/internal/services"

	"github.com/gofiber/fiber/v3"
)

type (
	PointLogsHandler interface {
		CreatePointLog(c *fiber.Ctx) error
	}

	pointLogsHandler struct {
		pointLogServices services.UserService
	}
)
