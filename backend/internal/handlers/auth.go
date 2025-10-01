package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *fiber.Ctx) error {
	var req SignupRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// TODO: Implement signup logic
	// - Validate input
	// - Hash password
	// - Save to database
	// - Generate JWT token

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user": fiber.Map{
			"name":  req.Name,
			"email": req.Email,
		},
	})
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// TODO: Implement login logic
	// - Validate input
	// - Find user in database
	// - Verify password
	// - Generate JWT token

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   "sample-jwt-token",
	})
}
