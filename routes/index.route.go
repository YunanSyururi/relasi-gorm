package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/controllers"
)

func RouteInit(app *fiber.App) {
	app.Get("/users", controllers.UserGetAll)
	app.Post("/users", controllers.CreateUser)

	app.Get("/lockers", controllers.LockerGetAll)
	app.Post("/lockers", controllers.CreateLocker)
}
