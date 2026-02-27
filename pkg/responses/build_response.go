package responses

import (
	"github.com/gofiber/fiber/v3"
)

func BuildErrorResponse(c fiber.Ctx, status int, message string, serviceCode string) error {
	return c.Status(status).JSON(Response{
		Status:      StatusMaps[status],
		StatusCode:  status,
		Message:     message,
		Response:    nil,
		ServiceCode: serviceCode,
	})
}

func BuildResponse(c fiber.Ctx, status int, message string, response interface{}, serviceCode string) error {
	return c.Status(status).JSON(Response{
		StatusCode:  status,
		Message:     message,
		Response:    response,
		ServiceCode: serviceCode,
	})
}
