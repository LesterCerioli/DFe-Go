package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Token represents the structure for a token
type Token struct {
	Value     string `json:"value" validate:"required"`
	IPAddress string `json:"ip_address" validate:"required"`
}

// Validate validates the Token fields
func (t *Token) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func main() {
	app := fiber.New()

	// Define the /token route
	app.Post("/token", func(c *fiber.Ctx) error {
		var tok Token
		// Parse the request body
		if err := c.BodyParser(&tok); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		// Validate the token
		if err := tok.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Handle token processing logic here
		return c.JSON(fiber.Map{"message": "Token received successfully", "token": tok})
	})

	// Start the Fiber app
	log.Fatal(app.Listen(":3000"))
}
