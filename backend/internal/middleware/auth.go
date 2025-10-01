package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Implement JWT authentication
		// - Get token from header
		// - Validate token
		// - Extract user info
		// - Set user in context

		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Validate token here

		return c.Next()
	}
}
