package authentication

import (
	"authenncrm/config"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c fiber.Ctx) error {
	cfg := config.ReadInConfig()

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}

		return []byte(cfg.Secret), nil
	})

	if err != nil || !token.Valid {

		if errors.Is(err, jwt.ErrTokenExpired) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token has expired",
			})
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	c.Locals("claims", claims)

	return c.Next()
}
