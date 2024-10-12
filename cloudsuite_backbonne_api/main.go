package main

import (
	"log"

	"cloudsuite_backbonne_api/handlers"
	"cloudsuite_backbonne_api/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	tokenService := services.NewTokenService("logs.json")

	tokenHandler := handlers.NewCreateTokenHandler(tokenService)

	app.Post("/token", tokenHandler.HandleTokenRequest)

	log.Fatal(app.Listen(":3000"))
}
