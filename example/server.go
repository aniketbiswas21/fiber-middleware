package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type ServerStatusStruct struct {
	Status  uint
	Message string
}

func main() {
	// initialise fiber app
	app := fiber.New()
	app.Use(logger.New())

	// landing page
	app.Get("/", func(c *fiber.Ctx) error {
		data := ServerStatusStruct{
			Status:  http.StatusOK,
			Message: "Server up and running on port 3000",
		}
		return c.JSON(data)
	})

	app.Listen(":3000")
}
