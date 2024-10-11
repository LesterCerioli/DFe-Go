package main

import (
	"cte_msc/database"
	"cte_msc/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	app.Get("/api/cte", handlers.GetCTEs(db))
	app.Post("/api/cte", handlers.CreateCTE(db))

	log.Fatal(app.Listen(":8080"))
}
