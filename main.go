package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/database"
	"github.com/muh-hizbe/relasi-gorm/database/migrations"
	"github.com/muh-hizbe/relasi-gorm/routes"
)

func main() {
	// Coonection to Database
	database.DatabaseInit()

	// MIGRATION
	migrations.Migration()

	// Fiber init
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		data := map[string]interface{}{
			"pesan": "ok",
			"usia":  24,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}

		return c.Send(jsonData)
	})
	routes.RouteInit(app)

	app.Listen(":8000")
}
