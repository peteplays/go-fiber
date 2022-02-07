package router

import (
	"go-fiber-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		err := ctx.SendString("And the API is UP!")
		return err
	})

	statusCheckRoutes := app.Group("/status")
	statusCheckRoutes.Get("/", controllers.GetHealthCheck)
	statusCheckRoutes.Get("/services", controllers.GetServiceCheck)

	swRoutes := app.Group("/sw")
	swRoutes.Get("/", controllers.GetSW)
	swRoutes.Get("/people/", controllers.GetPeople)
	swRoutes.Get("/people/:id", controllers.GetPerson)
}
