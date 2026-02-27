package responses

import "github.com/gofiber/fiber/v3"

var (
	StatusMaps = map[int]string{
		fiber.StatusBadRequest:          "Bad Request",
		fiber.StatusUnauthorized:        "Unauthorized",
		fiber.StatusForbidden:           "Forbidden",
		fiber.StatusNotFound:            "Not Found",
		fiber.StatusInternalServerError: "Internal Server Error",
		fiber.StatusNoContent:           "No Content",
		fiber.StatusCreated:             "Created",
		fiber.StatusOK:                  "OK",
		fiber.StatusUnprocessableEntity: "Unprocessable Entity",
	}
)
