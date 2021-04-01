package main

import (
	"fmt"
	"net/http"

	fiber_middleware "github.com/UniAuth/fiber-middleware"
	"github.com/UniAuth/fiber-middleware/models"
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

	var configs = []models.ApplicationConfig{
		models.ApplicationConfig{
			Name:            "server1",
			Url:             "http://localhost:5000",
			ClientId:        "600ee6ec924dd75267384cb4",
			ClientSecret:    "986727d0-c253-4adb-a9b8-c233a89cdb25",
			RedirectUri:     "http://localhost:3000/callback",
			AuthEndpoint:    "account/o/login",
			ProfileEndpoint: "account/o/access",
			ProfileProcessor: func(details string) {
				fmt.Println(details)
			},
		},
		models.ApplicationConfig{
			Name:            "server2",
			Url:             "http://localhost:4000",
			ClientId:        "600ee6ec924dd714752a8f74",
			ClientSecret:    "986727d0-147a-258a-3a6a-q7s4f5s8z3a6",
			RedirectUri:     "http://localhost:3000/callback",
			AuthEndpoint:    "account/o/login",
			ProfileEndpoint: "account/o/access",
			ProfileProcessor: func(details string) {
				fmt.Println(details)
			},
		},
	}

	// initialize uniAuth middleware instances
	uniAuth := fiber_middleware.Init(configs)

	// landing page
	app.Get("/", func(c *fiber.Ctx) error {
		data := ServerStatusStruct{
			Status:  http.StatusOK,
			Message: "Server up and running on port 3000",
		}
		return c.JSON(data)
	})

	// url to initiate login
	app.Get("/login", uniAuth.Authenticate("server1"))

	app.Listen(":3000")
}
