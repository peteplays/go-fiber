package main

import (
	"go-fiber-api/config"
	"go-fiber-api/router"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	config.Init()
}

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":" + os.Getenv("port"))
}
